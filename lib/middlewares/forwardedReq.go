package middlewares

import (
	apierrors "Proyect-Y/api-errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ForwardCustomHeader = "X-Gateway-Passed"


func DisAllowForwardedRequest(c *gin.Context) {
  _, ok := c.Request.Header[ForwardCustomHeader]

  if (ok) {
    c.JSON(http.StatusNotFound, apierrors.ServiceNotFound{
      Name: "NotFound",
      Message: "Path does not exists",
    })
    return
  }
  c.Next()
}
