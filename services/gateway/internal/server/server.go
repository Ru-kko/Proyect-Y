package server

import (
	"Proyect-Y/gateway/internal/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateServer(port int) {
	app := gin.Default()

	app.Any("/:service/*path", handlers.Gateway)

	app.Run(fmt.Sprintf(":%d", port))
}
