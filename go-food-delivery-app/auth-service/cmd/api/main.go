package main

import (
	"go-food-delivery-app/auth-service/internal/server"
	"go-food-delivery-app/auth-service/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	loadAllConfiguration()

	server.Start()
}

func loadAllConfiguration() {
	godotenv.Load(".env")

	server.LoadServerConfig()
	logger.LoadLoggerConfig()
}
