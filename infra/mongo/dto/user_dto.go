package dto

import "encoding/json"

type UserDto struct {
	Id       string `json:"id" bson:"_id,omitempty"`
	Email    string `json:"email" bson:"email"`
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
	Status   string `json:"status" bson:"status"`
}

func (o *UserDto) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}

func (o *UserDto) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, o)
}
