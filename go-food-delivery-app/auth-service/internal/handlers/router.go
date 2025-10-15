package handlers

import (
	"go-food-delivery-app/auth-service/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.Static("/front", "./front")

	// Initalize all services
	authService := services.NewAuthService()

	// Initalize all handlers
	authHandler := NewAuthHandler(authService)

	routerV1 := router.Group("api/v1")

	authRouter := routerV1.Group("auth")
	authRouter.POST("signup", authHandler.SignUpHandler)
	authRouter.POST("signin", authHandler.SignInHandler)

}
