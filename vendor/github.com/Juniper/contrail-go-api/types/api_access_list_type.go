//
// Automatically generated. DO NOT EDIT.
//

package types

type ApiAccessType struct {
	ApiName string `json:"api_name,omitempty"`
	Permissions *PermType `json:"permissions,omitempty"`
}

type ApiAccessListType struct {
	ApiAccess []ApiAccessType `json:"api_access,omitempty"`
}

func (obj *ApiAccessListType) AddApiAccess(value *ApiAccessType) {
        obj.ApiAccess = append(obj.ApiAccess, *value)
}
