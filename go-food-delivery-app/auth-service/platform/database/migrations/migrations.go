package migrations

import (
	_ "embed"
	"go-food-delivery-app/auth-service/internal/models"
	"go-food-delivery-app/auth-service/pkg/logger"
	"go-food-delivery-app/auth-service/platform/database"

	"go.uber.org/zap"
)

func ExecuteMigrations() (err error) {
	conf := database.NewDatabaseConfig()
	err = database.Connect(conf)
	if err != nil {
		logger.Log.Error("failed to connect to database",
			zap.Error(err),
		)
		return
	}

	// AutoMigrate the User model
	if err := database.DB.
		AutoMigrate(
			&models.UserCredentials{},
		); err != nil {
		logger.Log.Fatal("failed to auto-migrate",
			zap.Error(err),
		)
		return err
	}

	logger.Log.Info("Migration ran successfully")

	return nil
}
