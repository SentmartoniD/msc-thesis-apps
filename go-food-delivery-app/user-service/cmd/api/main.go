package main

import (
	"go-food-delivery-app/user-service/internal/platform/database"
	"go-food-delivery-app/user-service/internal/platform/database/migrations"
	"go-food-delivery-app/user-service/internal/server"
	"go-food-delivery-app/user-service/pkg/logger"
	"os"

	"go.uber.org/zap"
)

func main() {
	loadAllConfiguration()

	//connectToDatabase()

	server.Start()
}

func loadAllConfiguration() {
	server.LoadServerConfig()
	logger.LoadLoggerConfig()
}

func connectToDatabase() {
	databaseConfig := database.NewDatabaseConfig()
	err := database.Connect(databaseConfig)
	if err != nil {
		logger.Log.Error("failed to connect to the Database", zap.Error(err))
		os.Exit(1)
	}
	logger.Log.Info("Connected to the Database")

	err = migrations.ExecuteMigrations()
	if err != nil {
		logger.Log.Error("failed to execute migrations",
			zap.Error(err),
		)
		os.Exit(2)
	}
	logger.Log.Info("Migrations executed")
}
