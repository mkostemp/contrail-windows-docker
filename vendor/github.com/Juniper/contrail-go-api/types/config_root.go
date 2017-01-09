//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	config_root_api_access_list uint64 = 1 << iota
	config_root_id_perms
	config_root_display_name
	config_root_global_system_configs
	config_root_domains
)

type ConfigRoot struct {
        contrail.ObjectBase
	api_access_list ApiAccessListType
	id_perms IdPermsType
	display_name string
	global_system_configs contrail.ReferenceList
	domains contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *ConfigRoot) GetType() string {
        return "config-root"
}

func (obj *ConfigRoot) GetDefaultParent() []string {
        name := []string{}
        return name
}

func (obj *ConfigRoot) GetDefaultParentType() string {
        return ""
}

func (obj *ConfigRoot) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *ConfigRoot) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *ConfigRoot) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *ConfigRoot) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *ConfigRoot) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *ConfigRoot) GetApiAccessList() ApiAccessListType {
        return obj.api_access_list
}

func (obj *ConfigRoot) SetApiAccessList(value *ApiAccessListType) {
        obj.api_access_list = *value
        obj.modified |= config_root_api_access_list
}

func (obj *ConfigRoot) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *ConfigRoot) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= config_root_id_perms
}

func (obj *ConfigRoot) GetDisplayName() string {
        return obj.display_name
}

func (obj *ConfigRoot) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= config_root_display_name
}

func (obj *ConfigRoot) readGlobalSystemConfigs() error {
        if !obj.IsTransient() &&
                (obj.valid & config_root_global_system_configs == 0) {
                err := obj.GetField(obj, "global_system_configs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ConfigRoot) GetGlobalSystemConfigs() (
        contrail.ReferenceList, error) {
        err := obj.readGlobalSystemConfigs()
        if err != nil {
                return nil, err
        }
        return obj.global_system_configs, nil
}

func (obj *ConfigRoot) readDomains() error {
        if !obj.IsTransient() &&
                (obj.valid & config_root_domains == 0) {
                err := obj.GetField(obj, "domains")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ConfigRoot) GetDomains() (
        contrail.ReferenceList, error) {
        err := obj.readDomains()
        if err != nil {
                return nil, err
        }
        return obj.domains, nil
}

func (obj *ConfigRoot) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & config_root_api_access_list != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.api_access_list)
                if err != nil {
                        return nil, err
                }
                msg["api_access_list"] = &value
        }

        if obj.modified & config_root_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & config_root_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *ConfigRoot) UnmarshalJSON(body []byte) error {
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
                case "api_access_list":
                        err = json.Unmarshal(value, &obj.api_access_list)
                        if err == nil {
                                obj.valid |= config_root_api_access_list
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= config_root_id_perms
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= config_root_display_name
                        }
                        break
                case "global_system_configs":
                        err = json.Unmarshal(value, &obj.global_system_configs)
                        if err == nil {
                                obj.valid |= config_root_global_system_configs
                        }
                        break
                case "domains":
                        err = json.Unmarshal(value, &obj.domains)
                        if err == nil {
                                obj.valid |= config_root_domains
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ConfigRoot) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & config_root_api_access_list != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.api_access_list)
                if err != nil {
                        return nil, err
                }
                msg["api_access_list"] = &value
        }

        if obj.modified & config_root_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & config_root_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *ConfigRoot) UpdateReferences() error {

        return nil
}

func ConfigRootByName(c contrail.ApiClient, fqn string) (*ConfigRoot, error) {
    obj, err := c.FindByName("config-root", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*ConfigRoot), nil
}

func ConfigRootByUuid(c contrail.ApiClient, uuid string) (*ConfigRoot, error) {
    obj, err := c.FindByUuid("config-root", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*ConfigRoot), nil
}
