//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	service_instance_service_instance_properties uint64 = 1 << iota
	service_instance_id_perms
	service_instance_display_name
	service_instance_service_template_refs
	service_instance_virtual_machine_back_refs
	service_instance_logical_router_back_refs
	service_instance_loadbalancer_pool_back_refs
)

type ServiceInstance struct {
        contrail.ObjectBase
	service_instance_properties ServiceInstanceType
	id_perms IdPermsType
	display_name string
	service_template_refs contrail.ReferenceList
	virtual_machine_back_refs contrail.ReferenceList
	logical_router_back_refs contrail.ReferenceList
	loadbalancer_pool_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *ServiceInstance) GetType() string {
        return "service-instance"
}

func (obj *ServiceInstance) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *ServiceInstance) GetDefaultParentType() string {
        return "project"
}

func (obj *ServiceInstance) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *ServiceInstance) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *ServiceInstance) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *ServiceInstance) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *ServiceInstance) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *ServiceInstance) GetServiceInstanceProperties() ServiceInstanceType {
        return obj.service_instance_properties
}

func (obj *ServiceInstance) SetServiceInstanceProperties(value *ServiceInstanceType) {
        obj.service_instance_properties = *value
        obj.modified |= service_instance_service_instance_properties
}

func (obj *ServiceInstance) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *ServiceInstance) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= service_instance_id_perms
}

func (obj *ServiceInstance) GetDisplayName() string {
        return obj.display_name
}

func (obj *ServiceInstance) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= service_instance_display_name
}

func (obj *ServiceInstance) readServiceTemplateRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & service_instance_service_template_refs == 0) {
                err := obj.GetField(obj, "service_template_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceInstance) GetServiceTemplateRefs() (
        contrail.ReferenceList, error) {
        err := obj.readServiceTemplateRefs()
        if err != nil {
                return nil, err
        }
        return obj.service_template_refs, nil
}

func (obj *ServiceInstance) AddServiceTemplate(
        rhs *ServiceTemplate) error {
        err := obj.readServiceTemplateRefs()
        if err != nil {
                return err
        }

        if obj.modified & service_instance_service_template_refs == 0 {
                obj.storeReferenceBase("service-template", obj.service_template_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.service_template_refs = append(obj.service_template_refs, ref)
        obj.modified |= service_instance_service_template_refs
        return nil
}

func (obj *ServiceInstance) DeleteServiceTemplate(uuid string) error {
        err := obj.readServiceTemplateRefs()
        if err != nil {
                return err
        }

        if obj.modified & service_instance_service_template_refs == 0 {
                obj.storeReferenceBase("service-template", obj.service_template_refs)
        }

        for i, ref := range obj.service_template_refs {
                if ref.Uuid == uuid {
                        obj.service_template_refs = append(
                                obj.service_template_refs[:i],
                                obj.service_template_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= service_instance_service_template_refs
        return nil
}

func (obj *ServiceInstance) ClearServiceTemplate() {
        if (obj.valid & service_instance_service_template_refs != 0) &&
           (obj.modified & service_instance_service_template_refs == 0) {
                obj.storeReferenceBase("service-template", obj.service_template_refs)
        }
        obj.service_template_refs = make([]contrail.Reference, 0)
        obj.valid |= service_instance_service_template_refs
        obj.modified |= service_instance_service_template_refs
}

func (obj *ServiceInstance) SetServiceTemplateList(
        refList []contrail.ReferencePair) {
        obj.ClearServiceTemplate()
        obj.service_template_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.service_template_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *ServiceInstance) readVirtualMachineBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & service_instance_virtual_machine_back_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceInstance) GetVirtualMachineBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_back_refs, nil
}

func (obj *ServiceInstance) readLogicalRouterBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & service_instance_logical_router_back_refs == 0) {
                err := obj.GetField(obj, "logical_router_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceInstance) GetLogicalRouterBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readLogicalRouterBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.logical_router_back_refs, nil
}

func (obj *ServiceInstance) readLoadbalancerPoolBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & service_instance_loadbalancer_pool_back_refs == 0) {
                err := obj.GetField(obj, "loadbalancer_pool_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceInstance) GetLoadbalancerPoolBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readLoadbalancerPoolBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.loadbalancer_pool_back_refs, nil
}

func (obj *ServiceInstance) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & service_instance_service_instance_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_instance_properties)
                if err != nil {
                        return nil, err
                }
                msg["service_instance_properties"] = &value
        }

        if obj.modified & service_instance_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & service_instance_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.service_template_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_template_refs)
                if err != nil {
                        return nil, err
                }
                msg["service_template_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *ServiceInstance) UnmarshalJSON(body []byte) error {
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
                case "service_instance_properties":
                        err = json.Unmarshal(value, &obj.service_instance_properties)
                        if err == nil {
                                obj.valid |= service_instance_service_instance_properties
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= service_instance_id_perms
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= service_instance_display_name
                        }
                        break
                case "service_template_refs":
                        err = json.Unmarshal(value, &obj.service_template_refs)
                        if err == nil {
                                obj.valid |= service_instance_service_template_refs
                        }
                        break
                case "virtual_machine_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_back_refs)
                        if err == nil {
                                obj.valid |= service_instance_virtual_machine_back_refs
                        }
                        break
                case "logical_router_back_refs":
                        err = json.Unmarshal(value, &obj.logical_router_back_refs)
                        if err == nil {
                                obj.valid |= service_instance_logical_router_back_refs
                        }
                        break
                case "loadbalancer_pool_back_refs":
                        err = json.Unmarshal(value, &obj.loadbalancer_pool_back_refs)
                        if err == nil {
                                obj.valid |= service_instance_loadbalancer_pool_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceInstance) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & service_instance_service_instance_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_instance_properties)
                if err != nil {
                        return nil, err
                }
                msg["service_instance_properties"] = &value
        }

        if obj.modified & service_instance_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & service_instance_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & service_instance_service_template_refs != 0 {
                if len(obj.service_template_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["service_template_refs"] = &value
                } else if !obj.hasReferenceBase("service-template") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.service_template_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["service_template_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *ServiceInstance) UpdateReferences() error {

        if (obj.modified & service_instance_service_template_refs != 0) &&
           len(obj.service_template_refs) > 0 &&
           obj.hasReferenceBase("service-template") {
                err := obj.UpdateReference(
                        obj, "service-template",
                        obj.service_template_refs,
                        obj.baseMap["service-template"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func ServiceInstanceByName(c contrail.ApiClient, fqn string) (*ServiceInstance, error) {
    obj, err := c.FindByName("service-instance", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*ServiceInstance), nil
}

func ServiceInstanceByUuid(c contrail.ApiClient, uuid string) (*ServiceInstance, error) {
    obj, err := c.FindByUuid("service-instance", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*ServiceInstance), nil
}
