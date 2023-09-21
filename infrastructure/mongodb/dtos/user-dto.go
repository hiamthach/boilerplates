package dtos

import "encoding/json"

type UserDTO struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
}

func (u *UserDTO) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *UserDTO) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
