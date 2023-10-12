package typo

type ForwardedRequest[T any] struct {
	Data     T    `json:"data,omitempty" binding:"reuired"`
	AuthInfo Auth `json:"auth_info"`
}

type Auth struct {
	Id            string `json:"id,omitempty"`
	UserTag       string `json:"user_tag,omitempty" binding:"gte=4"`
	Authenticated bool   `json:"authenticated" binding:"required"`
	Roles         []Role `json:"roles,omitempty"`
	BornDate      string `json:"born_date,omitempty" bson:"born_date"`
}
