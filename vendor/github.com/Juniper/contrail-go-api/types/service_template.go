//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	service_template_service_template_properties uint64 = 1 << iota
	service_template_id_perms
	service_template_display_name
	service_template_service_instance_back_refs
)

type ServiceTemplate struct {
        contrail.ObjectBase
	service_template_properties ServiceTemplateType
	id_perms IdPermsType
	display_name string
	service_instance_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *ServiceTemplate) GetType() string {
        return "service-template"
}

func (obj *ServiceTemplate) GetDefaultParent() []string {
        name := []string{"default-domain"}
        return name
}

func (obj *ServiceTemplate) GetDefaultParentType() string {
        return "domain"
}

func (obj *ServiceTemplate) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *ServiceTemplate) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *ServiceTemplate) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *ServiceTemplate) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *ServiceTemplate) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *ServiceTemplate) GetServiceTemplateProperties() ServiceTemplateType {
        return obj.service_template_properties
}

func (obj *ServiceTemplate) SetServiceTemplateProperties(value *ServiceTemplateType) {
        obj.service_template_properties = *value
        obj.modified |= service_template_service_template_properties
}

func (obj *ServiceTemplate) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *ServiceTemplate) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= service_template_id_perms
}

func (obj *ServiceTemplate) GetDisplayName() string {
        return obj.display_name
}

func (obj *ServiceTemplate) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= service_template_display_name
}

func (obj *ServiceTemplate) readServiceInstanceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & service_template_service_instance_back_refs == 0) {
                err := obj.GetField(obj, "service_instance_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceTemplate) GetServiceInstanceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readServiceInstanceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.service_instance_back_refs, nil
}

func (obj *ServiceTemplate) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & service_template_service_template_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_template_properties)
                if err != nil {
                        return nil, err
                }
                msg["service_template_properties"] = &value
        }

        if obj.modified & service_template_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & service_template_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *ServiceTemplate) UnmarshalJSON(body []byte) error {
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
                case "service_template_properties":
                        err = json.Unmarshal(value, &obj.service_template_properties)
                        if err == nil {
                                obj.valid |= service_template_service_template_properties
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= service_template_id_perms
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= service_template_display_name
                        }
                        break
                case "service_instance_back_refs":
                        err = json.Unmarshal(value, &obj.service_instance_back_refs)
                        if err == nil {
                                obj.valid |= service_template_service_instance_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceTemplate) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & service_template_service_template_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_template_properties)
                if err != nil {
                        return nil, err
                }
                msg["service_template_properties"] = &value
        }

        if obj.modified & service_template_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & service_template_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *ServiceTemplate) UpdateReferences() error {

        return nil
}

func ServiceTemplateByName(c contrail.ApiClient, fqn string) (*ServiceTemplate, error) {
    obj, err := c.FindByName("service-template", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*ServiceTemplate), nil
}

func ServiceTemplateByUuid(c contrail.ApiClient, uuid string) (*ServiceTemplate, error) {
    obj, err := c.FindByUuid("service-template", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*ServiceTemplate), nil
}
