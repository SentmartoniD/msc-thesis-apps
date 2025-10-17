package main

import (
	"go-food-delivery-app/auth-service/internal/server"
	"go-food-delivery-app/auth-service/pkg/logger"
)

func main() {
	loadAllConfiguration()

	server.Start()
}

func loadAllConfiguration() {
	server.LoadServerConfig()
	logger.LoadLoggerConfig()
}
