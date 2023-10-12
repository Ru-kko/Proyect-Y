package handlers

import (
	apierrors "Proyect-Y/api-errors"
	"Proyect-Y/auth-service/internal/middleware"
	"Proyect-Y/typo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Indentify(c *gin.Context) {
	logger := logrus.New()
	token, isAuth := middleware.IsAuthenticated(c)

	if !isAuth {
		c.JSON(http.StatusOK, typo.Auth{
			Authenticated: false,
		})
		return
	}

	db, err := middleware.ReadDataInjection(c)

	if err != nil {
		logger.WithError(err).Error("Authentication Error: ")
		c.JSON(http.StatusInternalServerError, apierrors.InternalServerError{
			Name:    "InjectionError",
			Message: "We are having issues, try again later",
		})
		return
	}

	usr, err := db.GetUser(token.Id)

	if err != nil {
		logger.WithError(err).Error("Authentication Error: ")
		c.JSON(http.StatusInternalServerError, apierrors.InternalServerError{
			Name:    "DatabaseError",
			Message: "We are having issues, try again later",
		})
		return
	}

	c.JSON(http.StatusOK, typo.Auth{
		Id:            usr.Id,
		UserTag:       usr.UserTag,
		Authenticated: true,
		Roles:         usr.Roles,
		BornDate:      usr.BornDate,
	})
}
