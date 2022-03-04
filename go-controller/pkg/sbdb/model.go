// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package sbdb

import (
	"encoding/json"

	"github.com/ovn-org/libovsdb/model"
	"github.com/ovn-org/libovsdb/ovsdb"
)

// FullDatabaseModel returns the DatabaseModel object to be used in libovsdb
func FullDatabaseModel() (model.ClientDBModel, error) {
	return model.NewClientDBModel("OVN_Southbound", map[string]model.Model{
		"Address_Set":      &AddressSet{},
		"BFD":              &BFD{},
		"Chassis":          &Chassis{},
		"Chassis_Private":  &ChassisPrivate{},
		"Connection":       &Connection{},
		"Controller_Event": &ControllerEvent{},
		"DHCP_Options":     &DHCPOptions{},
		"DHCPv6_Options":   &DHCPv6Options{},
		"DNS":              &DNS{},
		"Datapath_Binding": &DatapathBinding{},
		"Encap":            &Encap{},
		"FDB":              &FDB{},
		"Gateway_Chassis":  &GatewayChassis{},
		"HA_Chassis":       &HAChassis{},
		"HA_Chassis_Group": &HAChassisGroup{},
		"IGMP_Group":       &IGMPGroup{},
		"IP_Multicast":     &IPMulticast{},
		"Load_Balancer":    &LoadBalancer{},
		"Logical_DP_Group": &LogicalDPGroup{},
		"Logical_Flow":     &LogicalFlow{},
		"MAC_Binding":      &MACBinding{},
		"Meter":            &Meter{},
		"Meter_Band":       &MeterBand{},
		"Multicast_Group":  &MulticastGroup{},
		"Port_Binding":     &PortBinding{},
		"Port_Group":       &PortGroup{},
		"RBAC_Permission":  &RBACPermission{},
		"RBAC_Role":        &RBACRole{},
		"SB_Global":        &SBGlobal{},
		"SSL":              &SSL{},
		"Service_Monitor":  &ServiceMonitor{},
	})
}

var schema = `{
  "name": "OVN_Southbound",
  "version": "20.21.0",
  "tables": {
    "Address_Set": {
      "columns": {
        "addresses": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "BFD": {
      "columns": {
        "detect_mult": {
          "type": "integer"
        },
        "disc": {
          "type": "integer"
        },
        "dst_ip": {
          "type": "string"
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "logical_port": {
          "type": "string"
        },
        "min_rx": {
          "type": "integer"
        },
        "min_tx": {
          "type": "integer"
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "src_port": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 49152,
              "maxInteger": 65535
            }
          }
        },
        "status": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "admin_down",
                  "down",
                  "init",
                  "up"
                ]
              ]
            }
          }
        }
      },
      "indexes": [
        [
          "logical_port",
          "dst_ip",
          "src_port",
          "disc"
        ]
      ]
    },
    "Chassis": {
      "columns": {
        "encaps": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Encap"
            },
            "max": "unlimited"
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "hostname": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "nb_cfg": {
          "type": "integer"
        },
        "other_config": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "transport_zones": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "vtep_logical_switches": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "Chassis_Private": {
      "columns": {
        "chassis": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Chassis",
              "refType": "weak"
            },
            "min": 0
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        },
        "nb_cfg": {
          "type": "integer"
        },
        "nb_cfg_timestamp": {
          "type": "integer"
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "Connection": {
      "columns": {
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "inactivity_probe": {
          "type": {
            "key": {
              "type": "integer"
            },
            "min": 0
          }
        },
        "is_connected": {
          "type": "boolean",
          "ephemeral": true
        },
        "max_backoff": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 1000
            },
            "min": 0
          }
        },
        "other_config": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "read_only": {
          "type": "boolean"
        },
        "role": {
          "type": "string"
        },
        "status": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          },
          "ephemeral": true
        },
        "target": {
          "type": "string"
        }
      },
      "indexes": [
        [
          "target"
        ]
      ]
    },
    "Controller_Event": {
      "columns": {
        "chassis": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Chassis",
              "refType": "weak"
            },
            "min": 0
          }
        },
        "event_info": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "event_type": {
          "type": {
            "key": {
              "type": "string",
              "enum": "empty_lb_backends"
            }
          }
        },
        "seq_num": {
          "type": "integer"
        }
      }
    },
    "DHCP_Options": {
      "columns": {
        "code": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 254
            }
          }
        },
        "name": {
          "type": "string"
        },
        "type": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "bool",
                  "domains",
                  "host_id",
                  "ipv4",
                  "static_routes",
                  "str",
                  "uint16",
                  "uint32",
                  "uint8"
                ]
              ]
            }
          }
        }
      }
    },
    "DHCPv6_Options": {
      "columns": {
        "code": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 254
            }
          }
        },
        "name": {
          "type": "string"
        },
        "type": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "ipv6",
                  "mac",
                  "str"
                ]
              ]
            }
          }
        }
      }
    },
    "DNS": {
      "columns": {
        "datapaths": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Datapath_Binding"
            },
            "max": "unlimited"
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "records": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      }
    },
    "Datapath_Binding": {
      "columns": {
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "load_balancers": {
          "type": {
            "key": {
              "type": "uuid"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "tunnel_key": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 1,
              "maxInteger": 16777215
            }
          }
        }
      },
      "indexes": [
        [
          "tunnel_key"
        ]
      ]
    },
    "Encap": {
      "columns": {
        "chassis_name": {
          "type": "string"
        },
        "ip": {
          "type": "string"
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "type": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "geneve",
                  "stt",
                  "vxlan"
                ]
              ]
            }
          }
        }
      },
      "indexes": [
        [
          "type",
          "ip"
        ]
      ]
    },
    "FDB": {
      "columns": {
        "dp_key": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 1,
              "maxInteger": 16777215
            }
          }
        },
        "mac": {
          "type": "string"
        },
        "port_key": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 1,
              "maxInteger": 16777215
            }
          }
        }
      },
      "indexes": [
        [
          "mac",
          "dp_key"
        ]
      ]
    },
    "Gateway_Chassis": {
      "columns": {
        "chassis": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Chassis",
              "refType": "weak"
            },
            "min": 0
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "priority": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 32767
            }
          }
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "HA_Chassis": {
      "columns": {
        "chassis": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Chassis",
              "refType": "weak"
            },
            "min": 0
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "priority": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 32767
            }
          }
        }
      }
    },
    "HA_Chassis_Group": {
      "columns": {
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ha_chassis": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "HA_Chassis"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        },
        "ref_chassis": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Chassis",
              "refType": "weak"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "IGMP_Group": {
      "columns": {
        "address": {
          "type": "string"
        },
        "chassis": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Chassis",
              "refType": "weak"
            },
            "min": 0
          }
        },
        "datapath": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Datapath_Binding",
              "refType": "weak"
            },
            "min": 0
          }
        },
        "ports": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Port_Binding",
              "refType": "weak"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      },
      "indexes": [
        [
          "address",
          "datapath",
          "chassis"
        ]
      ]
    },
    "IP_Multicast": {
      "columns": {
        "datapath": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Datapath_Binding",
              "refType": "weak"
            }
          }
        },
        "enabled": {
          "type": {
            "key": {
              "type": "boolean"
            },
            "min": 0
          }
        },
        "eth_src": {
          "type": "string"
        },
        "idle_timeout": {
          "type": {
            "key": {
              "type": "integer"
            },
            "min": 0
          }
        },
        "ip4_src": {
          "type": "string"
        },
        "ip6_src": {
          "type": "string"
        },
        "querier": {
          "type": {
            "key": {
              "type": "boolean"
            },
            "min": 0
          }
        },
        "query_interval": {
          "type": {
            "key": {
              "type": "integer"
            },
            "min": 0
          }
        },
        "query_max_resp": {
          "type": {
            "key": {
              "type": "integer"
            },
            "min": 0
          }
        },
        "seq_no": {
          "type": "integer"
        },
        "table_size": {
          "type": {
            "key": {
              "type": "integer"
            },
            "min": 0
          }
        }
      },
      "indexes": [
        [
          "datapath"
        ]
      ]
    },
    "Load_Balancer": {
      "columns": {
        "datapaths": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Datapath_Binding"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "protocol": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "sctp",
                  "tcp",
                  "udp"
                ]
              ]
            },
            "min": 0
          }
        },
        "vips": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      }
    },
    "Logical_DP_Group": {
      "columns": {
        "datapaths": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Datapath_Binding",
              "refType": "weak"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      }
    },
    "Logical_Flow": {
      "columns": {
        "actions": {
          "type": "string"
        },
        "controller_meter": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "logical_datapath": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Datapath_Binding"
            },
            "min": 0
          }
        },
        "logical_dp_group": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Logical_DP_Group"
            },
            "min": 0
          }
        },
        "match": {
          "type": "string"
        },
        "pipeline": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "egress",
                  "ingress"
                ]
              ]
            }
          }
        },
        "priority": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 65535
            }
          }
        },
        "table_id": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 32
            }
          }
        },
        "tags": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      }
    },
    "MAC_Binding": {
      "columns": {
        "datapath": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Datapath_Binding"
            }
          }
        },
        "ip": {
          "type": "string"
        },
        "logical_port": {
          "type": "string"
        },
        "mac": {
          "type": "string"
        }
      },
      "indexes": [
        [
          "logical_port",
          "ip"
        ]
      ]
    },
    "Meter": {
      "columns": {
        "bands": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Meter_Band"
            },
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        },
        "unit": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "kbps",
                  "pktps"
                ]
              ]
            }
          }
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "Meter_Band": {
      "columns": {
        "action": {
          "type": {
            "key": {
              "type": "string",
              "enum": "drop"
            }
          }
        },
        "burst_size": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 4294967295
            }
          }
        },
        "rate": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 1,
              "maxInteger": 4294967295
            }
          }
        }
      }
    },
    "Multicast_Group": {
      "columns": {
        "datapath": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Datapath_Binding"
            }
          }
        },
        "name": {
          "type": "string"
        },
        "ports": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Port_Binding",
              "refType": "weak"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "tunnel_key": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 32768,
              "maxInteger": 65535
            }
          }
        }
      },
      "indexes": [
        [
          "datapath",
          "tunnel_key"
        ],
        [
          "datapath",
          "name"
        ]
      ]
    },
    "Port_Binding": {
      "columns": {
        "chassis": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Chassis",
              "refType": "weak"
            },
            "min": 0
          }
        },
        "datapath": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Datapath_Binding"
            }
          }
        },
        "encap": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Encap",
              "refType": "weak"
            },
            "min": 0
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "gateway_chassis": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Gateway_Chassis"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ha_chassis_group": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "HA_Chassis_Group"
            },
            "min": 0
          }
        },
        "logical_port": {
          "type": "string"
        },
        "mac": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "nat_addresses": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "parent_port": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0
          }
        },
        "requested_chassis": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Chassis",
              "refType": "weak"
            },
            "min": 0
          }
        },
        "tag": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 1,
              "maxInteger": 4095
            },
            "min": 0
          }
        },
        "tunnel_key": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 1,
              "maxInteger": 32767
            }
          }
        },
        "type": {
          "type": "string"
        },
        "up": {
          "type": {
            "key": {
              "type": "boolean"
            },
            "min": 0
          }
        },
        "virtual_parent": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0
          }
        }
      },
      "indexes": [
        [
          "datapath",
          "tunnel_key"
        ],
        [
          "logical_port"
        ]
      ]
    },
    "Port_Group": {
      "columns": {
        "name": {
          "type": "string"
        },
        "ports": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "RBAC_Permission": {
      "columns": {
        "authorization": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "insert_delete": {
          "type": "boolean"
        },
        "table": {
          "type": "string"
        },
        "update": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      }
    },
    "RBAC_Role": {
      "columns": {
        "name": {
          "type": "string"
        },
        "permissions": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "uuid",
              "refTable": "RBAC_Permission",
              "refType": "weak"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      }
    },
    "SB_Global": {
      "columns": {
        "connections": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Connection"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ipsec": {
          "type": "boolean"
        },
        "nb_cfg": {
          "type": "integer"
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ssl": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "SSL"
            },
            "min": 0
          }
        }
      }
    },
    "SSL": {
      "columns": {
        "bootstrap_ca_cert": {
          "type": "boolean"
        },
        "ca_cert": {
          "type": "string"
        },
        "certificate": {
          "type": "string"
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "private_key": {
          "type": "string"
        },
        "ssl_ciphers": {
          "type": "string"
        },
        "ssl_protocols": {
          "type": "string"
        }
      }
    },
    "Service_Monitor": {
      "columns": {
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ip": {
          "type": "string"
        },
        "logical_port": {
          "type": "string"
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "port": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 65535
            }
          }
        },
        "protocol": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "tcp",
                  "udp"
                ]
              ]
            },
            "min": 0
          }
        },
        "src_ip": {
          "type": "string"
        },
        "src_mac": {
          "type": "string"
        },
        "status": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "error",
                  "offline",
                  "online"
                ]
              ]
            },
            "min": 0
          }
        }
      },
      "indexes": [
        [
          "logical_port",
          "ip",
          "port",
          "protocol"
        ]
      ]
    }
  }
}`

func Schema() ovsdb.DatabaseSchema {
	var s ovsdb.DatabaseSchema
	err := json.Unmarshal([]byte(schema), &s)
	if err != nil {
		panic(err)
	}
	return s
}