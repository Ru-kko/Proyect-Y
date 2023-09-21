package http

import (
	"Proyect-Y/gateway/internal/api/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateServer(port int) {
	app := gin.Default()

	app.Any(":/service/*path", routes.Gateway)

	app.Run(fmt.Sprintf(":%d", port))
}
