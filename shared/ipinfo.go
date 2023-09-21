// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    IPInfo, err := UnmarshalIPInfo(bytes)
//    bytes, err = IPInfo.Marshal()

package shared

import "encoding/json"

func UnmarshalIPInfo(data []byte) (IPInfo, error) {
	var r IPInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *IPInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type IPInfo struct {
	IP       *string `json:"ip,omitempty" bson:"ip,omitempty"`
	City     *string `json:"city,omitempty" bson:"city,omitempty"`
	Region   *string `json:"region,omitempty" bson:"region,omitempty"`
	Country  *string `json:"country,omitempty" bson:"country,omitempty"`
	LOC      *string `json:"loc,omitempty" bson:"loc,omitempty"`
	Org      *string `json:"org,omitempty" bson:"org,omitempty"`
	Postal   *string `json:"postal,omitempty" bson:"postal,omitempty"`
	Timezone *string `json:"timezone,omitempty" bson:"timezone,omitempty"`
}

func (o *IPInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *IPInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, o)
}
