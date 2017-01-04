//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	domain_domain_limits uint64 = 1 << iota
	domain_api_access_list
	domain_id_perms
	domain_display_name
	domain_projects
	domain_namespaces
	domain_service_templates
	domain_virtual_DNSs
)

type Domain struct {
        contrail.ObjectBase
	domain_limits DomainLimitsType
	api_access_list ApiAccessListType
	id_perms IdPermsType
	display_name string
	projects contrail.ReferenceList
	namespaces contrail.ReferenceList
	service_templates contrail.ReferenceList
	virtual_DNSs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *Domain) GetType() string {
        return "domain"
}

func (obj *Domain) GetDefaultParent() []string {
        name := []string{}
        return name
}

func (obj *Domain) GetDefaultParentType() string {
        return "config-root"
}

func (obj *Domain) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *Domain) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *Domain) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *Domain) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *Domain) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *Domain) GetDomainLimits() DomainLimitsType {
        return obj.domain_limits
}

func (obj *Domain) SetDomainLimits(value *DomainLimitsType) {
        obj.domain_limits = *value
        obj.modified |= domain_domain_limits
}

func (obj *Domain) GetApiAccessList() ApiAccessListType {
        return obj.api_access_list
}

func (obj *Domain) SetApiAccessList(value *ApiAccessListType) {
        obj.api_access_list = *value
        obj.modified |= domain_api_access_list
}

func (obj *Domain) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *Domain) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= domain_id_perms
}

func (obj *Domain) GetDisplayName() string {
        return obj.display_name
}

func (obj *Domain) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= domain_display_name
}

func (obj *Domain) readProjects() error {
        if !obj.IsTransient() &&
                (obj.valid & domain_projects == 0) {
                err := obj.GetField(obj, "projects")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Domain) GetProjects() (
        contrail.ReferenceList, error) {
        err := obj.readProjects()
        if err != nil {
                return nil, err
        }
        return obj.projects, nil
}

func (obj *Domain) readNamespaces() error {
        if !obj.IsTransient() &&
                (obj.valid & domain_namespaces == 0) {
                err := obj.GetField(obj, "namespaces")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Domain) GetNamespaces() (
        contrail.ReferenceList, error) {
        err := obj.readNamespaces()
        if err != nil {
                return nil, err
        }
        return obj.namespaces, nil
}

func (obj *Domain) readServiceTemplates() error {
        if !obj.IsTransient() &&
                (obj.valid & domain_service_templates == 0) {
                err := obj.GetField(obj, "service_templates")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Domain) GetServiceTemplates() (
        contrail.ReferenceList, error) {
        err := obj.readServiceTemplates()
        if err != nil {
                return nil, err
        }
        return obj.service_templates, nil
}

func (obj *Domain) readVirtualDnss() error {
        if !obj.IsTransient() &&
                (obj.valid & domain_virtual_DNSs == 0) {
                err := obj.GetField(obj, "virtual_DNSs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Domain) GetVirtualDnss() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualDnss()
        if err != nil {
                return nil, err
        }
        return obj.virtual_DNSs, nil
}

func (obj *Domain) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & domain_domain_limits != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.domain_limits)
                if err != nil {
                        return nil, err
                }
                msg["domain_limits"] = &value
        }

        if obj.modified & domain_api_access_list != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.api_access_list)
                if err != nil {
                        return nil, err
                }
                msg["api_access_list"] = &value
        }

        if obj.modified & domain_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & domain_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *Domain) UnmarshalJSON(body []byte) error {
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
                case "domain_limits":
                        err = json.Unmarshal(value, &obj.domain_limits)
                        if err == nil {
                                obj.valid |= domain_domain_limits
                        }
                        break
                case "api_access_list":
                        err = json.Unmarshal(value, &obj.api_access_list)
                        if err == nil {
                                obj.valid |= domain_api_access_list
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= domain_id_perms
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= domain_display_name
                        }
                        break
                case "projects":
                        err = json.Unmarshal(value, &obj.projects)
                        if err == nil {
                                obj.valid |= domain_projects
                        }
                        break
                case "namespaces":
                        err = json.Unmarshal(value, &obj.namespaces)
                        if err == nil {
                                obj.valid |= domain_namespaces
                        }
                        break
                case "service_templates":
                        err = json.Unmarshal(value, &obj.service_templates)
                        if err == nil {
                                obj.valid |= domain_service_templates
                        }
                        break
                case "virtual_DNSs":
                        err = json.Unmarshal(value, &obj.virtual_DNSs)
                        if err == nil {
                                obj.valid |= domain_virtual_DNSs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Domain) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & domain_domain_limits != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.domain_limits)
                if err != nil {
                        return nil, err
                }
                msg["domain_limits"] = &value
        }

        if obj.modified & domain_api_access_list != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.api_access_list)
                if err != nil {
                        return nil, err
                }
                msg["api_access_list"] = &value
        }

        if obj.modified & domain_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & domain_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *Domain) UpdateReferences() error {

        return nil
}

func DomainByName(c contrail.ApiClient, fqn string) (*Domain, error) {
    obj, err := c.FindByName("domain", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*Domain), nil
}

func DomainByUuid(c contrail.ApiClient, uuid string) (*Domain, error) {
    obj, err := c.FindByUuid("domain", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*Domain), nil
}
