package handlers

import (
	apierrors "Proyect-Y/api-errors"
	"Proyect-Y/auth-service/internal/domain"
	"Proyect-Y/auth-service/internal/middleware"
	"Proyect-Y/auth-service/internal/security"
	"Proyect-Y/typo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type tokenResponse struct {
	Type       string `json:"toke_type"`
	Token      string `json:"token"`
	Expiration int64  `json:"expiration"`
	UserID     string `json:"user_id"`
	UserTag    string `json:"user_tag"`
}

func SignIn(c *gin.Context) {
	var reqBody typo.ForwardedRequest[domain.AuthCredentials]
	logger := logrus.New()

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, apierrors.BadRequest{
			Name:    "BadRequest",
			Message: "Bad information",
		})
		return
	}

	service, err := middleware.ReadDataInjection(c)
	if err != nil {
		logger.WithError(err).Error("Singin error: ")
		c.JSON(http.StatusInternalServerError, apierrors.InternalServerError{
			Name: "InjectionError",
			Code: 500,
		})
	}

	usr, err := service.GetUserByTag(reqBody.Data.UserTag)
	if err != nil {
		logger.WithError(err).Error("Singin error: ")
		c.JSON(http.StatusInternalServerError, apierrors.InternalServerError{
			Name: "DataBaseError",
			Code: 500,
		})
		return
	}

	if usr == nil {
		c.JSON(http.StatusNotFound, apierrors.UserNotFound{
			Name: "UserNotFound",
			User: reqBody.Data.UserTag,
		})
		return
	}

	err = security.VerifyPassword(reqBody.Data.Password, usr.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, apierrors.NotAuthorizedError{
			Name:    "PasswordNotMatch",
			Message: "Incorret password",
		})
		return
	}

	token, exp, err := security.BuildToken(*usr)
	if err != nil {
		logger.WithError(err).Error("Singin error: ")
		c.JSON(http.StatusInternalServerError, apierrors.InternalServerError{
			Code:    500,
			Message: "Couldnot generate token",
			Name:    "EncryptError",
		})
		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		Type: "Bearer",
		Token: token,
		Expiration: exp,
		UserID: usr.Id,
		UserTag: usr.UserTag,
	})
}
