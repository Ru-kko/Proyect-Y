package security

import (
	"Proyect-Y/auth-service/internal/domain"
	"Proyect-Y/auth-service/internal/util"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func BuildToken(data domain.StoredUser) (string, int64, error) {
	env := util.GetEnv()
	now := time.Now().Unix()
	expirationTime := now + (60 * 60 * 24 * 12) // 12 days

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTclaims{
		AuthenticatedInfo: domain.AuthenticatedInfo {
			Id: data.Id,
			UserTag: data.UserTag,
		},
		Exp:        expirationTime,
		Iss:        now,
	})

	token, err := jwt.SignedString([]byte(env.JWT_SECRET))
	return token, expirationTime, err
}

func ValidateToken(token string) (*JWTclaims, error) {
	env := util.GetEnv()
	claims := &JWTclaims{}

	res, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) {
		return []byte(env.JWT_SECRET), nil
	})

	if err != nil || !res.Valid {
		return nil, err
	}

	claims, ok := res.Claims.(*JWTclaims)
	if !ok {
		return nil, &JWTError{
			Type: "Fromat Error",
			Message: "Could not parse token to type JWTclaims",
		}
	}

	return claims, err
}

