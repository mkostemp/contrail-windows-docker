//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	global_vrouter_config_linklocal_services uint64 = 1 << iota
	global_vrouter_config_encapsulation_priorities
	global_vrouter_config_vxlan_network_identifier_mode
	global_vrouter_config_flow_export_rate
	global_vrouter_config_flow_aging_timeout_list
	global_vrouter_config_forwarding_mode
	global_vrouter_config_id_perms
	global_vrouter_config_display_name
)

type GlobalVrouterConfig struct {
        contrail.ObjectBase
	linklocal_services LinklocalServicesTypes
	encapsulation_priorities EncapsulationPrioritiesType
	vxlan_network_identifier_mode string
	flow_export_rate int
	flow_aging_timeout_list FlowAgingTimeoutList
	forwarding_mode string
	id_perms IdPermsType
	display_name string
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *GlobalVrouterConfig) GetType() string {
        return "global-vrouter-config"
}

func (obj *GlobalVrouterConfig) GetDefaultParent() []string {
        name := []string{"default-global-system-config"}
        return name
}

func (obj *GlobalVrouterConfig) GetDefaultParentType() string {
        return "global-system-config"
}

func (obj *GlobalVrouterConfig) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *GlobalVrouterConfig) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *GlobalVrouterConfig) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *GlobalVrouterConfig) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *GlobalVrouterConfig) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *GlobalVrouterConfig) GetLinklocalServices() LinklocalServicesTypes {
        return obj.linklocal_services
}

func (obj *GlobalVrouterConfig) SetLinklocalServices(value *LinklocalServicesTypes) {
        obj.linklocal_services = *value
        obj.modified |= global_vrouter_config_linklocal_services
}

func (obj *GlobalVrouterConfig) GetEncapsulationPriorities() EncapsulationPrioritiesType {
        return obj.encapsulation_priorities
}

func (obj *GlobalVrouterConfig) SetEncapsulationPriorities(value *EncapsulationPrioritiesType) {
        obj.encapsulation_priorities = *value
        obj.modified |= global_vrouter_config_encapsulation_priorities
}

func (obj *GlobalVrouterConfig) GetVxlanNetworkIdentifierMode() string {
        return obj.vxlan_network_identifier_mode
}

func (obj *GlobalVrouterConfig) SetVxlanNetworkIdentifierMode(value string) {
        obj.vxlan_network_identifier_mode = value
        obj.modified |= global_vrouter_config_vxlan_network_identifier_mode
}

func (obj *GlobalVrouterConfig) GetFlowExportRate() int {
        return obj.flow_export_rate
}

func (obj *GlobalVrouterConfig) SetFlowExportRate(value int) {
        obj.flow_export_rate = value
        obj.modified |= global_vrouter_config_flow_export_rate
}

func (obj *GlobalVrouterConfig) GetFlowAgingTimeoutList() FlowAgingTimeoutList {
        return obj.flow_aging_timeout_list
}

func (obj *GlobalVrouterConfig) SetFlowAgingTimeoutList(value *FlowAgingTimeoutList) {
        obj.flow_aging_timeout_list = *value
        obj.modified |= global_vrouter_config_flow_aging_timeout_list
}

func (obj *GlobalVrouterConfig) GetForwardingMode() string {
        return obj.forwarding_mode
}

func (obj *GlobalVrouterConfig) SetForwardingMode(value string) {
        obj.forwarding_mode = value
        obj.modified |= global_vrouter_config_forwarding_mode
}

func (obj *GlobalVrouterConfig) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *GlobalVrouterConfig) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= global_vrouter_config_id_perms
}

func (obj *GlobalVrouterConfig) GetDisplayName() string {
        return obj.display_name
}

func (obj *GlobalVrouterConfig) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= global_vrouter_config_display_name
}

func (obj *GlobalVrouterConfig) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & global_vrouter_config_linklocal_services != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.linklocal_services)
                if err != nil {
                        return nil, err
                }
                msg["linklocal_services"] = &value
        }

        if obj.modified & global_vrouter_config_encapsulation_priorities != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.encapsulation_priorities)
                if err != nil {
                        return nil, err
                }
                msg["encapsulation_priorities"] = &value
        }

        if obj.modified & global_vrouter_config_vxlan_network_identifier_mode != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.vxlan_network_identifier_mode)
                if err != nil {
                        return nil, err
                }
                msg["vxlan_network_identifier_mode"] = &value
        }

        if obj.modified & global_vrouter_config_flow_export_rate != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.flow_export_rate)
                if err != nil {
                        return nil, err
                }
                msg["flow_export_rate"] = &value
        }

        if obj.modified & global_vrouter_config_flow_aging_timeout_list != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.flow_aging_timeout_list)
                if err != nil {
                        return nil, err
                }
                msg["flow_aging_timeout_list"] = &value
        }

        if obj.modified & global_vrouter_config_forwarding_mode != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.forwarding_mode)
                if err != nil {
                        return nil, err
                }
                msg["forwarding_mode"] = &value
        }

        if obj.modified & global_vrouter_config_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & global_vrouter_config_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *GlobalVrouterConfig) UnmarshalJSON(body []byte) error {
        var m map[string]json.RawMessage
        err := json.Unmarshal(body, &m)
        if err != nil {
                return err
        }
        err = obj.UnmarshalCommon(m)
        if err != nil {
                return err
        }
        for key, value := range m {
                switch key {
                case "linklocal_services":
                        err = json.Unmarshal(value, &obj.linklocal_services)
                        if err == nil {
                                obj.valid |= global_vrouter_config_linklocal_services
                        }
                        break
                case "encapsulation_priorities":
                        err = json.Unmarshal(value, &obj.encapsulation_priorities)
                        if err == nil {
                                obj.valid |= global_vrouter_config_encapsulation_priorities
                        }
                        break
                case "vxlan_network_identifier_mode":
                        err = json.Unmarshal(value, &obj.vxlan_network_identifier_mode)
                        if err == nil {
                                obj.valid |= global_vrouter_config_vxlan_network_identifier_mode
                        }
                        break
                case "flow_export_rate":
                        err = json.Unmarshal(value, &obj.flow_export_rate)
                        if err == nil {
                                obj.valid |= global_vrouter_config_flow_export_rate
                        }
                        break
                case "flow_aging_timeout_list":
                        err = json.Unmarshal(value, &obj.flow_aging_timeout_list)
                        if err == nil {
                                obj.valid |= global_vrouter_config_flow_aging_timeout_list
                        }
                        break
                case "forwarding_mode":
                        err = json.Unmarshal(value, &obj.forwarding_mode)
                        if err == nil {
                                obj.valid |= global_vrouter_config_forwarding_mode
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= global_vrouter_config_id_perms
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= global_vrouter_config_display_name
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *GlobalVrouterConfig) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & global_vrouter_config_linklocal_services != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.linklocal_services)
                if err != nil {
                        return nil, err
                }
                msg["linklocal_services"] = &value
        }

        if obj.modified & global_vrouter_config_encapsulation_priorities != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.encapsulation_priorities)
                if err != nil {
                        return nil, err
                }
                msg["encapsulation_priorities"] = &value
        }

        if obj.modified & global_vrouter_config_vxlan_network_identifier_mode != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.vxlan_network_identifier_mode)
                if err != nil {
                        return nil, err
                }
                msg["vxlan_network_identifier_mode"] = &value
        }

        if obj.modified & global_vrouter_config_flow_export_rate != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.flow_export_rate)
                if err != nil {
                        return nil, err
                }
                msg["flow_export_rate"] = &value
        }

        if obj.modified & global_vrouter_config_flow_aging_timeout_list != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.flow_aging_timeout_list)
                if err != nil {
                        return nil, err
                }
                msg["flow_aging_timeout_list"] = &value
        }

        if obj.modified & global_vrouter_config_forwarding_mode != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.forwarding_mode)
                if err != nil {
                        return nil, err
                }
                msg["forwarding_mode"] = &value
        }

        if obj.modified & global_vrouter_config_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & global_vrouter_config_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *GlobalVrouterConfig) UpdateReferences() error {

        return nil
}

func GlobalVrouterConfigByName(c contrail.ApiClient, fqn string) (*GlobalVrouterConfig, error) {
    obj, err := c.FindByName("global-vrouter-config", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*GlobalVrouterConfig), nil
}

func GlobalVrouterConfigByUuid(c contrail.ApiClient, uuid string) (*GlobalVrouterConfig, error) {
    obj, err := c.FindByUuid("global-vrouter-config", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*GlobalVrouterConfig), nil
}
