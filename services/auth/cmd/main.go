package main

import (
	"os"

	"Proyect-Y/auth-service/internal/server"
	"Proyect-Y/auth-service/internal/util"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("PROD") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}
	env := util.GetEnv()

	server.CreateServer(env.PORT)
}
