// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package nbdb

import "github.com/ovn-org/libovsdb/model"

// PortGroup defines an object in Port_Group table
type PortGroup struct {
	UUID        string            `ovsdb:"_uuid"`
	ACLs        []string          `ovsdb:"acls"`
	ExternalIDs map[string]string `ovsdb:"external_ids"`
	Name        string            `ovsdb:"name"`
	Ports       []string          `ovsdb:"ports"`
}

func copyPortGroupACLs(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalPortGroupACLs(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func copyPortGroupExternalIDs(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalPortGroupExternalIDs(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func copyPortGroupPorts(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalPortGroupPorts(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func (a *PortGroup) DeepCopyInto(b *PortGroup) {
	*b = *a
	b.ACLs = copyPortGroupACLs(a.ACLs)
	b.ExternalIDs = copyPortGroupExternalIDs(a.ExternalIDs)
	b.Ports = copyPortGroupPorts(a.Ports)
}

func (a *PortGroup) DeepCopy() *PortGroup {
	b := new(PortGroup)
	a.DeepCopyInto(b)
	return b
}

func (a *PortGroup) CloneModelInto(b model.Model) {
	c := b.(*PortGroup)
	a.DeepCopyInto(c)
}

func (a *PortGroup) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *PortGroup) Equals(b *PortGroup) bool {
	return a.UUID == b.UUID &&
		equalPortGroupACLs(a.ACLs, b.ACLs) &&
		equalPortGroupExternalIDs(a.ExternalIDs, b.ExternalIDs) &&
		a.Name == b.Name &&
		equalPortGroupPorts(a.Ports, b.Ports)
}

func (a *PortGroup) EqualsModel(b model.Model) bool {
	c := b.(*PortGroup)
	return a.Equals(c)
}

var _ model.CloneableModel = &PortGroup{}
var _ model.ComparableModel = &PortGroup{}