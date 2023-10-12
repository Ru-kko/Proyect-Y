package domain

type AuthCredentials struct {
	UserTag  string `json:"user_tag" binding:"required,gte=4"`
	Password string `json:"password" binding:"required"`
}

type AuthenticatedInfo struct {
	Id       string    `json:"id,omitempty"`
	UserTag  string    `json:"user_tag"`
}
