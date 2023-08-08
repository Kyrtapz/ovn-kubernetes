package ovnwebhook

import (
	"context"
	"fmt"

	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/types"
	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/util"
	"golang.org/x/exp/maps"
	corev1 "k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type checkNodeAnnot func(v annotationChange, nodeName string) error

// commonNodeAnnotationChecks holds annotations allowed for ovnkube-node:<nodeName> users in non-IC and IC environments
var commonNodeAnnotationChecks = map[string]checkNodeAnnot{
	util.OvnNodeHostAddresses:            nil,
	util.OvnNodeL3GatewayConfig:          nil,
	util.OvnNodeManagementPortMacAddress: nil,
	util.OvnNodeIfAddr:                   nil,
	util.OvnNodeChassisID: func(v annotationChange, nodeName string) error {
		if v.action == removed {
			return fmt.Errorf("%s cannot be removed", util.OvnNodeChassisID)
		}
		if v.action == changed {
			return fmt.Errorf("%s cannot be changed once set", util.OvnNodeChassisID)
		}
		return nil
	},
	util.OvnNodeZoneName: func(v annotationChange, nodeName string) error {
		// it is allowed for the annotation to be set to "global" or <nodeName> initially
		if v.action == added && (v.value == types.OvnDefaultZone || v.value == nodeName) {
			return nil
		}

		// it is allowed for the annotation to be changed to <nodeName>
		if v.action == changed && v.value == nodeName {
			return nil
		}

		return fmt.Errorf("%s can only be set to %s or %s, it cannot be removed", util.OvnNodeZoneName, types.OvnDefaultZone, nodeName)
	},
}

// interconnectNodeAnnotationChecks holds annotations allowed for ovnkube-node:<nodeName> users in IC environments
var interconnectNodeAnnotationChecks = map[string]checkNodeAnnot{
	util.OvnNodeMigratedZoneName: func(v annotationChange, nodeName string) error {
		// it is allowed for the annotation to be set to <nodeName>
		if (v.action == added || v.action == changed) && v.value == nodeName {
			return nil
		}

		return fmt.Errorf("%s can only be set to %s, it cannot be removed", util.OvnNodeMigratedZoneName, nodeName)
	},
}

type NodeAdmission struct {
	annotationChecks map[string]checkNodeAnnot
	annotationKeys   sets.Set[string]
}

func NewNodeAdmissionWebhook(enableInterconnect bool) *NodeAdmission {
	checks := commonNodeAnnotationChecks
	if enableInterconnect {
		maps.Copy(checks, interconnectNodeAnnotationChecks)
	}
	return &NodeAdmission{
		annotationChecks: checks,
		annotationKeys:   sets.New[string](maps.Keys(checks)...),
	}
}

var _ admission.CustomValidator = &NodeAdmission{}

func (p NodeAdmission) ValidateCreate(_ context.Context, _ runtime.Object) (warnings admission.Warnings, err error) {
	// Ignore creation, the webhook is configured to only handle nodes/status updates
	return nil, nil
}

func (p NodeAdmission) ValidateDelete(_ context.Context, _ runtime.Object) (warnings admission.Warnings, err error) {
	// Ignore deletion, the webhook is configured to only handle nodes/status updates
	return nil, nil
}

func (p NodeAdmission) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (warnings admission.Warnings, err error) {
	oldNode := oldObj.(*corev1.Node)
	newNode := newObj.(*corev1.Node)

	changes := mapDiff(oldNode.Annotations, newNode.Annotations)
	changedKeys := maps.Keys(changes)

	req, err := admission.RequestFromContext(ctx)
	if err != nil {
		return nil, err
	}
	nodeName, isOVNKubeNode := ovnkubeNodeIdentity(req.UserInfo)

	if len(changes) == 0 && !isOVNKubeNode {
		// no annotation changes and the user is not ovnkube-node
		return nil, nil
	}

	if !isOVNKubeNode {
		// the user is not an ovnkube-node
		return nil, nil
	}

	if newNode.Name != nodeName {
		return nil, fmt.Errorf("ovnkube-node on node: %q is not allowed to modify nodes %q annotations", nodeName, newNode.Name)
	}

	// ovnkube-node is not allowed to change annotations outside of it's scope
	if !p.annotationKeys.HasAll(changedKeys...) {
		return nil, fmt.Errorf("ovnkube-node on node: %q is not allowed to set at least one of the annotations: %v", nodeName, changedKeys)
	}

	for _, key := range changedKeys {
		if check := p.annotationChecks[key]; check != nil {
			if err := check(changes[key], nodeName); err != nil {
				return nil, fmt.Errorf("ovnkube-node on node: %q is not allowed to set %s on node %q: %v", nodeName, key, newNode.Name, err)
			}
		}
	}

	// Verify that nothing but the annotations changed.
	// Since ovnkube-node only has the node/status permissions, it is enough to check .Status and .ObjectMeta only.
	// Ignore .ManagedFields fields which are modified on every update.
	oldNodeShallowCopy := oldNode
	newNodeShallowCopy := newNode
	oldNodeShallowCopy.Annotations = nil
	newNodeShallowCopy.Annotations = nil
	oldNodeShallowCopy.ManagedFields = nil
	newNodeShallowCopy.ManagedFields = nil
	if !apiequality.Semantic.DeepEqual(oldNodeShallowCopy.ObjectMeta, newNodeShallowCopy.ObjectMeta) ||
		!apiequality.Semantic.DeepEqual(oldNodeShallowCopy.Status, newNodeShallowCopy.Status) {
		return nil, fmt.Errorf("ovnkube-node on node: %q is not allowed to modify anything other than annotations", nodeName)
	}

	return nil, nil
}