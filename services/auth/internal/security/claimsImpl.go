package security

import (
	"Proyect-Y/auth-service/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTclaims struct {
	domain.StoredUser
	Iss int64
	Exp int64
}

func (jw *JWTclaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: time.Unix(jw.Exp, 0),
	}, nil
}

func (jw *JWTclaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: time.Unix(jw.Iss, 0),
	}, nil
}

func (jw *JWTclaims) GetNotBefore() (*jwt.NumericDate, error) {
	return jw.GetIssuedAt()
}

func (jw *JWTclaims) GetIssuer() (string, error) {
	return "AUTH", nil
}

func (jw *JWTclaims) GetSubject() (string, error) {
	return jw.UserTag, nil
}

func (jw *JWTclaims) GetAudience() (jwt.ClaimStrings, error) {
	return []string{"APP"}, nil
}
