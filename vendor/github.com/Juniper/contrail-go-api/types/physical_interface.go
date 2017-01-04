//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	physical_interface_id_perms uint64 = 1 << iota
	physical_interface_display_name
	physical_interface_logical_interfaces
)

type PhysicalInterface struct {
        contrail.ObjectBase
	id_perms IdPermsType
	display_name string
	logical_interfaces contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *PhysicalInterface) GetType() string {
        return "physical-interface"
}

func (obj *PhysicalInterface) GetDefaultParent() []string {
        name := []string{"default-global-system-config", "default-physical-router"}
        return name
}

func (obj *PhysicalInterface) GetDefaultParentType() string {
        return "physical-router"
}

func (obj *PhysicalInterface) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *PhysicalInterface) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *PhysicalInterface) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *PhysicalInterface) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *PhysicalInterface) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *PhysicalInterface) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *PhysicalInterface) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= physical_interface_id_perms
}

func (obj *PhysicalInterface) GetDisplayName() string {
        return obj.display_name
}

func (obj *PhysicalInterface) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= physical_interface_display_name
}

func (obj *PhysicalInterface) readLogicalInterfaces() error {
        if !obj.IsTransient() &&
                (obj.valid & physical_interface_logical_interfaces == 0) {
                err := obj.GetField(obj, "logical_interfaces")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalInterface) GetLogicalInterfaces() (
        contrail.ReferenceList, error) {
        err := obj.readLogicalInterfaces()
        if err != nil {
                return nil, err
        }
        return obj.logical_interfaces, nil
}

func (obj *PhysicalInterface) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & physical_interface_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & physical_interface_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *PhysicalInterface) UnmarshalJSON(body []byte) error {
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
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= physical_interface_id_perms
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= physical_interface_display_name
                        }
                        break
                case "logical_interfaces":
                        err = json.Unmarshal(value, &obj.logical_interfaces)
                        if err == nil {
                                obj.valid |= physical_interface_logical_interfaces
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalInterface) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & physical_interface_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & physical_interface_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *PhysicalInterface) UpdateReferences() error {

        return nil
}

func PhysicalInterfaceByName(c contrail.ApiClient, fqn string) (*PhysicalInterface, error) {
    obj, err := c.FindByName("physical-interface", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*PhysicalInterface), nil
}

func PhysicalInterfaceByUuid(c contrail.ApiClient, uuid string) (*PhysicalInterface, error) {
    obj, err := c.FindByUuid("physical-interface", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*PhysicalInterface), nil
}
