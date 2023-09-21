package main

import (
	"os"
	"strconv"

	"Proyect-Y/auth-service/internal/util"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("PROD") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()

	
	app.Run(":" + strconv.Itoa(util.GetEnv().PORT))
}
