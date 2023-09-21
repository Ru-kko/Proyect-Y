package main

import (
	"Proyect-Y/gateway/internal/api/http"
	"Proyect-Y/gateway/internal/util"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("PROD") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	http.CreateServer(util.GetEnv().PORT)
}
