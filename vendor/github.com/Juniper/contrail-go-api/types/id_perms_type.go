//
// Automatically generated. DO NOT EDIT.
//

package types

type UuidType struct {
	UuidMslong uint64 `json:"uuid_mslong,omitempty"`
	UuidLslong uint64 `json:"uuid_lslong,omitempty"`
}

type IdPermsType struct {
	Permissions *PermType `json:"permissions,omitempty"`
	Uuid *UuidType `json:"uuid,omitempty"`
	Enable bool `json:"enable,omitempty"`
	Created string `json:"created,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
	Description string `json:"description,omitempty"`
	UserVisible bool `json:"user_visible,omitempty"`
	Creator string `json:"creator,omitempty"`
}
