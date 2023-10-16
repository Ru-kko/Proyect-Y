package main

import (
	"Proyect-Y/gateway/internal/server"
	"Proyect-Y/gateway/internal/util"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("PROD") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	server.CreateServer(util.GetEnv().PORT)
}
