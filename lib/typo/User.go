package typo

import "Proyect-Y/typo/constants/locations"

type User struct {
	Id        string             `json:"id,omitempty" bson:"_id,omitempty"`
	UserTag   string             `json:"user_tag" bson:"user_tag"`
	BornDate  string             `json:"born_date" bson:"born_date"`
	Country   locations.Location `json:"country" bson:"country"`
	FristName string             `json:"frist_name" bson:"frist_name"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
}
