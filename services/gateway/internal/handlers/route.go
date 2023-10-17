package handlers

import (
	apierrors "Proyect-Y/api-errors"
	serv "Proyect-Y/gateway/internal/service"
	"Proyect-Y/gateway/internal/util"
	"Proyect-Y/middlewares"
	"Proyect-Y/typo"
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	l "github.com/sirupsen/logrus"
)

func Gateway(c *gin.Context) {
	logger := l.New()
	service := c.Param("service")
	path := c.Param("path")
	method := c.Request.Method

	logger.WithFields(l.Fields{
		"service": service,
		"path":    path,
		"method":  method,
	})

	host, err := util.GetService(service)
	if err != nil {
		switch e := err.(type) {
		case *apierrors.ServiceNotFound:
			c.JSON(http.StatusNotFound, e)
			return
		}
	}

	query := c.Request.URL.RawQuery
	url := util.BuildUrl(host, path, query)

	auth, err := serv.GetAuthentication(c.GetHeader("Authorization"))
	if err != nil {
		logger.WithError(err).Log(l.ErrorLevel)
		c.JSON(http.StatusInternalServerError, apierrors.InternalServerError{
			Name:    "AuthenticationError",
			Code:    500,
			Message: "We having issues in authntication service",
		})
		return
	}

	var reqBody any = ""
	if c.Request.ContentLength > 0 {
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, apierrors.BadRequest{
				Name:    "BadRequest",
				Message: "Bad information",
			})
			return
		}
	}

	body := typo.ForwardedRequest[any]{
		Data:     reqBody,
		AuthInfo: *auth,
	}
	rawStr, err := json.Marshal(body)
	if err != nil {
		logger.WithError(err).Log(l.ErrorLevel)
		c.JSON(http.StatusInternalServerError, apierrors.InternalServerError{
			Name: "ParseError",
			Code: 500,
		})
		return
	}

	req, err := http.NewRequest(c.Request.Method, url, bytes.NewBuffer(rawStr))
	if err != nil {
		logger.WithError(err).Log(l.ErrorLevel)
		c.JSON(http.StatusInternalServerError, apierrors.InternalServerError{
			Name: "GatewayError",
			Code: 500,
		})
	}

	req.Header = c.Request.Header
	req.Header.Add(middlewares.ForwardCustomHeader, "true")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "redirect error"})
		logger.WithError(err).Log(l.ErrorLevel)
		return
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Body reading error"})
		logger.WithError(err).Log(l.ErrorLevel)
		return
	}

	c.Writer.WriteHeader(res.StatusCode)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Write(resBody)
}
