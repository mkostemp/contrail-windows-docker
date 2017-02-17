//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	bridge_domain_mac_learning_enabled uint64 = 1 << iota
	bridge_domain_mac_limit_control
	bridge_domain_mac_move_control
	bridge_domain_mac_aging_time
	bridge_domain_isid
	bridge_domain_id_perms
	bridge_domain_perms2
	bridge_domain_annotations
	bridge_domain_display_name
	bridge_domain_virtual_machine_interface_back_refs
)

type BridgeDomain struct {
        contrail.ObjectBase
	mac_learning_enabled bool
	mac_limit_control MACLimitControlType
	mac_move_control MACMoveLimitControlType
	mac_aging_time int
	isid int
	id_perms IdPermsType
	perms2 PermType2
	annotations KeyValuePairs
	display_name string
	virtual_machine_interface_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *BridgeDomain) GetType() string {
        return "bridge-domain"
}

func (obj *BridgeDomain) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project", "default-virtual-network"}
        return name
}

func (obj *BridgeDomain) GetDefaultParentType() string {
        return "virtual-network"
}

func (obj *BridgeDomain) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *BridgeDomain) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *BridgeDomain) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *BridgeDomain) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *BridgeDomain) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *BridgeDomain) GetMacLearningEnabled() bool {
        return obj.mac_learning_enabled
}

func (obj *BridgeDomain) SetMacLearningEnabled(value bool) {
        obj.mac_learning_enabled = value
        obj.modified |= bridge_domain_mac_learning_enabled
}

func (obj *BridgeDomain) GetMacLimitControl() MACLimitControlType {
        return obj.mac_limit_control
}

func (obj *BridgeDomain) SetMacLimitControl(value *MACLimitControlType) {
        obj.mac_limit_control = *value
        obj.modified |= bridge_domain_mac_limit_control
}

func (obj *BridgeDomain) GetMacMoveControl() MACMoveLimitControlType {
        return obj.mac_move_control
}

func (obj *BridgeDomain) SetMacMoveControl(value *MACMoveLimitControlType) {
        obj.mac_move_control = *value
        obj.modified |= bridge_domain_mac_move_control
}

func (obj *BridgeDomain) GetMacAgingTime() int {
        return obj.mac_aging_time
}

func (obj *BridgeDomain) SetMacAgingTime(value int) {
        obj.mac_aging_time = value
        obj.modified |= bridge_domain_mac_aging_time
}

func (obj *BridgeDomain) GetIsid() int {
        return obj.isid
}

func (obj *BridgeDomain) SetIsid(value int) {
        obj.isid = value
        obj.modified |= bridge_domain_isid
}

func (obj *BridgeDomain) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *BridgeDomain) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= bridge_domain_id_perms
}

func (obj *BridgeDomain) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *BridgeDomain) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= bridge_domain_perms2
}

func (obj *BridgeDomain) GetAnnotations() KeyValuePairs {
        return obj.annotations
}

func (obj *BridgeDomain) SetAnnotations(value *KeyValuePairs) {
        obj.annotations = *value
        obj.modified |= bridge_domain_annotations
}

func (obj *BridgeDomain) GetDisplayName() string {
        return obj.display_name
}

func (obj *BridgeDomain) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= bridge_domain_display_name
}

func (obj *BridgeDomain) readVirtualMachineInterfaceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & bridge_domain_virtual_machine_interface_back_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *BridgeDomain) GetVirtualMachineInterfaceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_back_refs, nil
}

func (obj *BridgeDomain) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & bridge_domain_mac_learning_enabled != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.mac_learning_enabled)
                if err != nil {
                        return nil, err
                }
                msg["mac_learning_enabled"] = &value
        }

        if obj.modified & bridge_domain_mac_limit_control != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.mac_limit_control)
                if err != nil {
                        return nil, err
                }
                msg["mac_limit_control"] = &value
        }

        if obj.modified & bridge_domain_mac_move_control != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.mac_move_control)
                if err != nil {
                        return nil, err
                }
                msg["mac_move_control"] = &value
        }

        if obj.modified & bridge_domain_mac_aging_time != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.mac_aging_time)
                if err != nil {
                        return nil, err
                }
                msg["mac_aging_time"] = &value
        }

        if obj.modified & bridge_domain_isid != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.isid)
                if err != nil {
                        return nil, err
                }
                msg["isid"] = &value
        }

        if obj.modified & bridge_domain_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & bridge_domain_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & bridge_domain_annotations != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified & bridge_domain_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *BridgeDomain) UnmarshalJSON(body []byte) error {
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
                case "mac_learning_enabled":
                        err = json.Unmarshal(value, &obj.mac_learning_enabled)
                        if err == nil {
                                obj.valid |= bridge_domain_mac_learning_enabled
                        }
                        break
                case "mac_limit_control":
                        err = json.Unmarshal(value, &obj.mac_limit_control)
                        if err == nil {
                                obj.valid |= bridge_domain_mac_limit_control
                        }
                        break
                case "mac_move_control":
                        err = json.Unmarshal(value, &obj.mac_move_control)
                        if err == nil {
                                obj.valid |= bridge_domain_mac_move_control
                        }
                        break
                case "mac_aging_time":
                        err = json.Unmarshal(value, &obj.mac_aging_time)
                        if err == nil {
                                obj.valid |= bridge_domain_mac_aging_time
                        }
                        break
                case "isid":
                        err = json.Unmarshal(value, &obj.isid)
                        if err == nil {
                                obj.valid |= bridge_domain_isid
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= bridge_domain_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= bridge_domain_perms2
                        }
                        break
                case "annotations":
                        err = json.Unmarshal(value, &obj.annotations)
                        if err == nil {
                                obj.valid |= bridge_domain_annotations
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= bridge_domain_display_name
                        }
                        break
                case "virtual_machine_interface_back_refs": {
                        type ReferenceElement struct {
                                To []string
                                Uuid string
                                Href string
                                Attr BridgeDomainMembershipType
                        }
                        var array []ReferenceElement
                        err = json.Unmarshal(value, &array)
                        if err != nil {
                            break
                        }
                        obj.valid |= bridge_domain_virtual_machine_interface_back_refs
                        obj.virtual_machine_interface_back_refs = make(contrail.ReferenceList, 0)
                        for _, element := range array {
                                ref := contrail.Reference {
                                        element.To,
                                        element.Uuid,
                                        element.Href,
                                        element.Attr,
                                }
                                obj.virtual_machine_interface_back_refs = append(obj.virtual_machine_interface_back_refs, ref)
                        }
                        break
                }
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *BridgeDomain) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & bridge_domain_mac_learning_enabled != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.mac_learning_enabled)
                if err != nil {
                        return nil, err
                }
                msg["mac_learning_enabled"] = &value
        }

        if obj.modified & bridge_domain_mac_limit_control != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.mac_limit_control)
                if err != nil {
                        return nil, err
                }
                msg["mac_limit_control"] = &value
        }

        if obj.modified & bridge_domain_mac_move_control != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.mac_move_control)
                if err != nil {
                        return nil, err
                }
                msg["mac_move_control"] = &value
        }

        if obj.modified & bridge_domain_mac_aging_time != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.mac_aging_time)
                if err != nil {
                        return nil, err
                }
                msg["mac_aging_time"] = &value
        }

        if obj.modified & bridge_domain_isid != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.isid)
                if err != nil {
                        return nil, err
                }
                msg["isid"] = &value
        }

        if obj.modified & bridge_domain_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & bridge_domain_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & bridge_domain_annotations != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified & bridge_domain_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *BridgeDomain) UpdateReferences() error {

        return nil
}

func BridgeDomainByName(c contrail.ApiClient, fqn string) (*BridgeDomain, error) {
    obj, err := c.FindByName("bridge-domain", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*BridgeDomain), nil
}

func BridgeDomainByUuid(c contrail.ApiClient, uuid string) (*BridgeDomain, error) {
    obj, err := c.FindByUuid("bridge-domain", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*BridgeDomain), nil
}
