package typo

type ForwardedRequest[T any] struct {
	Data     T    `json:"data,omitempty" binding:"required"`
	AuthInfo Auth `json:"auth_info" binding:"omitempty"`
}

type Auth struct {
	Id            string `json:"id,omitempty"`
	UserTag       string `json:"user_tag,omitempty" binding:"omitempty,gte=4"`
	Authenticated bool   `json:"authenticated" binding:"required"`
	Roles         []Role `json:"roles,omitempty"`
	BornDate      string `json:"born_date,omitempty"`
}
