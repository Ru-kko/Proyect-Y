package typo

import "Proyect-Y/typo/constants/locations"

type AuthData struct {
	Id       string `json:"id,omitempty" bson:"_id,omitempty"`
	UserTag  string `json:"user_tag" bson:"user_tag"`
	BornDate string `json:"born_date" bson:"born_date"`
	Password string `json:"passwrod" bson:"passwrod"`
	Roles    role   `json:"roles" bson:"roles"`
	Email    string `json:"email" bson:"email"`
}

type RegisterData struct {
	UserTag   string             `json:"user_tag" bson:"user_tag"`
	BornDate  string             `json:"born_date" bson:"born_date"`
	Password  string             `json:"passwrod" bson:"passwrod"`
	Email     string             `json:"email" bson:"email"`
	Country   locations.Location `json:"country" bson:"country"`
	FristName string             `json:"frist_name" bson:"frist_name"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
}

type role string

const (
	Admin_Rol role = "admin"
	User_Rol  role = "user"
)
