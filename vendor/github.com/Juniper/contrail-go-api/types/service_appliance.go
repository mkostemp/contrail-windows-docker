//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	service_appliance_service_appliance_user_credentials uint64 = 1 << iota
	service_appliance_service_appliance_ip_address
	service_appliance_service_appliance_properties
	service_appliance_id_perms
	service_appliance_display_name
)

type ServiceAppliance struct {
        contrail.ObjectBase
	service_appliance_user_credentials UserCredentials
	service_appliance_ip_address string
	service_appliance_properties KeyValuePairs
	id_perms IdPermsType
	display_name string
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *ServiceAppliance) GetType() string {
        return "service-appliance"
}

func (obj *ServiceAppliance) GetDefaultParent() []string {
        name := []string{"default-global-system-config", "default-service-appliance-set"}
        return name
}

func (obj *ServiceAppliance) GetDefaultParentType() string {
        return "service-appliance-set"
}

func (obj *ServiceAppliance) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *ServiceAppliance) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *ServiceAppliance) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *ServiceAppliance) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *ServiceAppliance) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *ServiceAppliance) GetServiceApplianceUserCredentials() UserCredentials {
        return obj.service_appliance_user_credentials
}

func (obj *ServiceAppliance) SetServiceApplianceUserCredentials(value *UserCredentials) {
        obj.service_appliance_user_credentials = *value
        obj.modified |= service_appliance_service_appliance_user_credentials
}

func (obj *ServiceAppliance) GetServiceApplianceIpAddress() string {
        return obj.service_appliance_ip_address
}

func (obj *ServiceAppliance) SetServiceApplianceIpAddress(value string) {
        obj.service_appliance_ip_address = value
        obj.modified |= service_appliance_service_appliance_ip_address
}

func (obj *ServiceAppliance) GetServiceApplianceProperties() KeyValuePairs {
        return obj.service_appliance_properties
}

func (obj *ServiceAppliance) SetServiceApplianceProperties(value *KeyValuePairs) {
        obj.service_appliance_properties = *value
        obj.modified |= service_appliance_service_appliance_properties
}

func (obj *ServiceAppliance) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *ServiceAppliance) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= service_appliance_id_perms
}

func (obj *ServiceAppliance) GetDisplayName() string {
        return obj.display_name
}

func (obj *ServiceAppliance) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= service_appliance_display_name
}

func (obj *ServiceAppliance) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & service_appliance_service_appliance_user_credentials != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_user_credentials)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_user_credentials"] = &value
        }

        if obj.modified & service_appliance_service_appliance_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_ip_address"] = &value
        }

        if obj.modified & service_appliance_service_appliance_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_properties)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_properties"] = &value
        }

        if obj.modified & service_appliance_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & service_appliance_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *ServiceAppliance) UnmarshalJSON(body []byte) error {
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
                case "service_appliance_user_credentials":
                        err = json.Unmarshal(value, &obj.service_appliance_user_credentials)
                        if err == nil {
                                obj.valid |= service_appliance_service_appliance_user_credentials
                        }
                        break
                case "service_appliance_ip_address":
                        err = json.Unmarshal(value, &obj.service_appliance_ip_address)
                        if err == nil {
                                obj.valid |= service_appliance_service_appliance_ip_address
                        }
                        break
                case "service_appliance_properties":
                        err = json.Unmarshal(value, &obj.service_appliance_properties)
                        if err == nil {
                                obj.valid |= service_appliance_service_appliance_properties
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= service_appliance_id_perms
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= service_appliance_display_name
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceAppliance) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & service_appliance_service_appliance_user_credentials != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_user_credentials)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_user_credentials"] = &value
        }

        if obj.modified & service_appliance_service_appliance_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_ip_address"] = &value
        }

        if obj.modified & service_appliance_service_appliance_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_properties)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_properties"] = &value
        }

        if obj.modified & service_appliance_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & service_appliance_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *ServiceAppliance) UpdateReferences() error {

        return nil
}

func ServiceApplianceByName(c contrail.ApiClient, fqn string) (*ServiceAppliance, error) {
    obj, err := c.FindByName("service-appliance", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*ServiceAppliance), nil
}

func ServiceApplianceByUuid(c contrail.ApiClient, uuid string) (*ServiceAppliance, error) {
    obj, err := c.FindByUuid("service-appliance", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*ServiceAppliance), nil
}
