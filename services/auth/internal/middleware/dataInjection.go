package middleware

import (
	apierrors "Proyect-Y/api-errors"
	"Proyect-Y/auth-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const dataServiceName = "DataService"

func DataServiceInject() gin.HandlerFunc {
  logger := logrus.New()
	return func(c *gin.Context) {
		injectable, err := service.NewDataService()

		if err != nil {
      logger.WithError(err).Error("Injection error: ")
			c.JSON(http.StatusInternalServerError, apierrors.InternalServerError{Name: "InternalServerError", Code: 500})
			return
		}

    c.Set(dataServiceName, injectable)
    c.Next()

    injectable.CloseAll()
	}
}

func ReadDataInjection(c *gin.Context) (*service.DataService, error) {
  err := apierrors.InternalServerError{Code: 500, Name: "Bad Injection"}

  dataService, exists := c.Get(dataServiceName)
	if !exists {
    return nil, err
	}

  data, ok := dataService.(*service.DataService)
  if !ok {
    return nil, err
  }

  return data, nil
}
