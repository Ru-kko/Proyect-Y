package routes

import (
	apierrors "Proyect-Y/api-errors"
	"Proyect-Y/gateway/internal/util"
	"bytes"
	"fmt"
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
			c.JSON(http.StatusNotFound, gin.H{"message": e.Name})
			break
		default:
			logger.WithError(err).Log(l.ErrorLevel)
			break
		}
		return
	}

	query := c.Request.URL.RawQuery
	url := fmt.Sprintf("%s%s", host, path)

	if query != "" {
		url = fmt.Sprintf("%s?%s", url, query)
	}
	reqBody, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	req, _ := http.NewRequest(c.Request.Method, url, bytes.NewReader(reqBody))
	req.Header = c.Request.Header
	req.Header.Add("X-OUT-COMMING", "true")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "redirect error"})
		logger.WithError(err).Log(l.ErrorLevel)
		return
	}
	defer res.Body.Close()

	for key, values := range res.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Body reading error"})
		logger.WithError(err).Log(l.ErrorLevel)
		return
	}

	c.String(res.StatusCode, string(body))
}
