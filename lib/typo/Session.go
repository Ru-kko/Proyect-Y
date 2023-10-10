package typo

import "Proyect-Y/typo/constants/locations"

type AuthChange struct {
	Id       string `json:"id,omitempty" bson:"_id,omitempty" binding:"required"`
	UserTag  string `json:"user_tag" bson:"user_tag" binding:"gte=4"`
	BornDate string `json:"born_date" bson:"born_date"`
}

type AuthenticationInfo struct {
	Id            string `json:"id,omitempty"`
	UserTag       string `json:"user_tag,omitempty"`
	Authenticated bool   `json:"authenticated"`
	Roles         []Role `json:"roles" binding:"required"`
}

type RegisterData struct {
	UserTag   string             `json:"user_tag" bson:"user_tag" binding:"required,gte=4"`
	BornDate  string             `json:"born_date" bson:"born_date" binding:"required"`
	Password  string             `json:"password" bson:"password" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required,email"`
	Country   locations.Location `json:"country" bson:"country" binding:"required"`
	FirstName string             `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty" `
}

type Role string

const (
	No_Auth   Role = "unauth"
	Admin_Rol Role = "admin"
	User_Rol  Role = "user"
)
