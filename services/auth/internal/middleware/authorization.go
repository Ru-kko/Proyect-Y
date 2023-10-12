package middleware

import (
	apierrors "Proyect-Y/api-errors"
	"Proyect-Y/auth-service/internal/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const TokenScope = "Token"


func AutorizeMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.Next()
		return
	}

	token, err := security.ValidateToken(strings.Replace(auth, "Bearer ", "", 1))
	if err != nil {
		c.JSON(http.StatusUnauthorized, apierrors.NotAuthorizedError{
			Name:    "BadToken",
			Message: "Invalid token",
		})
		return
	}

	c.Set(TokenScope, token)
	c.Next()
}

func IsAuthenticated(c *gin.Context) (*security.JWTclaims, bool) {
	auth, exists := c.Get(TokenScope)

	if !exists {
		return nil, false
	}

	val, ok := auth.(*security.JWTclaims)
	logrus.Info(val, ", ok:", ok, ", exists:", exists, ", auth:", auth)

	return val, ok
}
