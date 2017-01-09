//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	interface_route_table_interface_route_table_routes uint64 = 1 << iota
	interface_route_table_id_perms
	interface_route_table_display_name
	interface_route_table_virtual_machine_interface_back_refs
)

type InterfaceRouteTable struct {
        contrail.ObjectBase
	interface_route_table_routes RouteTableType
	id_perms IdPermsType
	display_name string
	virtual_machine_interface_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *InterfaceRouteTable) GetType() string {
        return "interface-route-table"
}

func (obj *InterfaceRouteTable) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *InterfaceRouteTable) GetDefaultParentType() string {
        return "project"
}

func (obj *InterfaceRouteTable) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *InterfaceRouteTable) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *InterfaceRouteTable) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *InterfaceRouteTable) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *InterfaceRouteTable) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *InterfaceRouteTable) GetInterfaceRouteTableRoutes() RouteTableType {
        return obj.interface_route_table_routes
}

func (obj *InterfaceRouteTable) SetInterfaceRouteTableRoutes(value *RouteTableType) {
        obj.interface_route_table_routes = *value
        obj.modified |= interface_route_table_interface_route_table_routes
}

func (obj *InterfaceRouteTable) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *InterfaceRouteTable) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= interface_route_table_id_perms
}

func (obj *InterfaceRouteTable) GetDisplayName() string {
        return obj.display_name
}

func (obj *InterfaceRouteTable) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= interface_route_table_display_name
}

func (obj *InterfaceRouteTable) readVirtualMachineInterfaceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & interface_route_table_virtual_machine_interface_back_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *InterfaceRouteTable) GetVirtualMachineInterfaceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_back_refs, nil
}

func (obj *InterfaceRouteTable) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & interface_route_table_interface_route_table_routes != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.interface_route_table_routes)
                if err != nil {
                        return nil, err
                }
                msg["interface_route_table_routes"] = &value
        }

        if obj.modified & interface_route_table_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & interface_route_table_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *InterfaceRouteTable) UnmarshalJSON(body []byte) error {
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
                case "interface_route_table_routes":
                        err = json.Unmarshal(value, &obj.interface_route_table_routes)
                        if err == nil {
                                obj.valid |= interface_route_table_interface_route_table_routes
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= interface_route_table_id_perms
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= interface_route_table_display_name
                        }
                        break
                case "virtual_machine_interface_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
                        if err == nil {
                                obj.valid |= interface_route_table_virtual_machine_interface_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *InterfaceRouteTable) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & interface_route_table_interface_route_table_routes != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.interface_route_table_routes)
                if err != nil {
                        return nil, err
                }
                msg["interface_route_table_routes"] = &value
        }

        if obj.modified & interface_route_table_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & interface_route_table_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *InterfaceRouteTable) UpdateReferences() error {

        return nil
}

func InterfaceRouteTableByName(c contrail.ApiClient, fqn string) (*InterfaceRouteTable, error) {
    obj, err := c.FindByName("interface-route-table", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*InterfaceRouteTable), nil
}

func InterfaceRouteTableByUuid(c contrail.ApiClient, uuid string) (*InterfaceRouteTable, error) {
    obj, err := c.FindByUuid("interface-route-table", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*InterfaceRouteTable), nil
}
