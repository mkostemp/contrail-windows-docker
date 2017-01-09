//
// Automatically generated. DO NOT EDIT.
//

package types

type AllowedAddressPair struct {
	Ip *SubnetType `json:"ip,omitempty"`
	Mac string `json:"mac,omitempty"`
	AddressMode string `json:"address_mode,omitempty"`
}

type AllowedAddressPairs struct {
	AllowedAddressPair []AllowedAddressPair `json:"allowed_address_pair,omitempty"`
}

func (obj *AllowedAddressPairs) AddAllowedAddressPair(value *AllowedAddressPair) {
        obj.AllowedAddressPair = append(obj.AllowedAddressPair, *value)
}
