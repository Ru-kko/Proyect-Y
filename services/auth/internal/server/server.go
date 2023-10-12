package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateServer(port int) {
  app := gin.Default()

  getAuthRoutes(app)
  getUserRoutes(app)

  app.Run(fmt.Sprintf(":%d", port))
}
