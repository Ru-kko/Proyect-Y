package typo

import "Proyect-Y/typo/constants/locations"

type AuthData struct {
	Id       string `json:"id,omitempty" bson:"_id,omitempty"`
	UserTag  string `json:"user_tag" bson:"user_tag" binding:"required,gte=4"`
	BornDate string `json:"born_date" bson:"born_date" binding:"required"`
	Password string `json:"passwrod" bson:"passwrod" binding:"required"`
	Roles    role   `json:"roles" bson:"roles"`
	Email    string `json:"email" bson:"email" binding:"required,email"`
}

type AuthChange struct {
	Id       string `json:"id,omitempty" bson:"_id,omitempty" binding:"required"`
	UserTag  string `json:"user_tag" bson:"user_tag" binding:"gte=4"`
	BornDate string `json:"born_date" bson:"born_date"`
}

type AuthCredentials struct {
	UserTag  string `json:"user_tag" binding:"required,gte=4"`
	Password string `json:"passwrod" binding:"required"`
}

type RegisterData struct {
	UserTag   string             `json:"user_tag" bson:"user_tag" binding:"required,gte=4"`
	BornDate  string             `json:"born_date" bson:"born_date" binding:"required"`
	Password  string             `json:"passwrod" bson:"passwrod" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required,email"`
	Country   locations.Location `json:"country" bson:"country" binding:"required"`
	FristName string             `json:"frist_name" bson:"frist_name" binding:"required"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty" `
}

type role string

const (
	Admin_Rol role = "admin"
	User_Rol  role = "user"
)
