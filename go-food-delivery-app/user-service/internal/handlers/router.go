package handlers

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	router.Static("/front", "./front")
}
