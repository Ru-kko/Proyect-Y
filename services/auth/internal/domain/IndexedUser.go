package domain

import "Proyect-Y/typo"

type StoredUser struct {
	Id       string    `json:"id,omitempty" bson:"_id,omitempty"`
	UserTag  string    `json:"user_tag" bson:"user_tag" binding:"required,gte=4"`
	Password string    `json:"password" bson:"password" binding:"required"`
	BornDate string    `json:"born_date" bson:"born_date" binding:"required"`
	Roles    typo.Role `json:"roles" bson:"roles"`
	Email    string    `json:"email" bson:"email" binding:"required,email"`
}
