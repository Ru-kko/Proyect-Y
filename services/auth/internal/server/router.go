package server

import (
	"Proyect-Y/auth-service/internal/handlers"
	"Proyect-Y/auth-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func getAuthRoutes(e *gin.Engine) {
  router := e.Group("/")

  router.Use(middleware.DataServiceInject())
  router.POST("/signin", handlers.SignIn)
  router.POST("/signup", handlers.SignUp)
}
