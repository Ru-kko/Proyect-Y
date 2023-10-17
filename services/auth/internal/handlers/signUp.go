package handlers

import (
	apierrors "Proyect-Y/api-errors"
	"Proyect-Y/auth-service/internal/middleware"
	"Proyect-Y/auth-service/internal/security"
	"Proyect-Y/typo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func SignUp(c *gin.Context) {
	var reqBody typo.ForwardedRequest[typo.RegisterData]
	logger := logrus.New()

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, apierrors.BadRequest{
			Name:    "BadRequest",
			Message: err.Error(),
		})
		return
	}

	service, err := middleware.ReadDataInjection(c)
	if err != nil {
		logger.WithError(err).Error("singup error: ")
		c.JSON(http.StatusInternalServerError, apierrors.InternalServerError{
			Name: "InjectionError",
			Code: 500,
		})
		return
	}

	user, err := service.UserRegister(reqBody.Data)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteError); ok && writeErr.HasErrorCode(11000) {
			c.JSON(http.StatusBadRequest, apierrors.BadRequest{
				Name:    "BadRequest",
				Message: "Repeated usertag or email",
			})
			return
		}
		logger.WithError(err).Error("singup error: ")
		c.JSON(http.StatusInternalServerError, apierrors.InternalServerError{
			Name: "InternalServerError",
			Code: 500,
		})
		return
	}

	token, exp, err := security.BuildToken(*user)
	if err != nil {
		logger.WithError(err).Error("singup error: ")
		c.JSON(http.StatusInternalServerError, apierrors.InternalServerError{
			Name:    "SessionError",
			Code:    500,
			Message: "Couldn't create session token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "id": user.Id, "type": "Bearer", "exp": exp, "user": user.UserTag})
}

