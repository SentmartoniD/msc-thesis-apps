package migrations

import (
	_ "embed"
	"go-food-delivery-app/user-service/internal/platform/database"
	"go-food-delivery-app/user-service/pkg/logger"

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
		//&models.User{},
		); err != nil {
		logger.Log.Fatal("failed to auto-migrate",
			zap.Error(err),
		)
		return err
	}

	// var existingAdmin models.User
	// check := database.DB.Find(&existingAdmin)
	// if check.Error != nil {
	// 	err = check.Error
	// 	return
	// }

	// if check.RowsAffected == 0 {

	// 	fakeCard := models.Card{
	// 		ID:          uuid.NewString(),
	// 		Number:      "1234567890",
	// 		IsValidated: true,
	// 	}

	// 	encryptedPassword := crypto.HashPassword("5tysX6B6.3!?")

	// 	admins := []models.User{
	// 		{
	// 			FirstName: "admin",
	// 			LastName:  "admin",
	// 			Email:     "info@bvsbservices.ch",
	// 			Password:  string(encryptedPassword),
	// 			Picture:   "", Role: enums.ADMIN,
	// 			IsActive: true,
	// 			Card:     fakeCard,
	// 		},
	// 		{
	// 			FirstName: "admin",
	// 			LastName:  "admin",
	// 			Email:     "miljan.zekovic1@gmail.com",
	// 			Password:  string(encryptedPassword),
	// 			Picture:   "", Role: enums.ADMIN,
	// 			IsActive: true,
	// 			Card:     fakeCard,
	// 		},
	// 	}

	// 	for _, admin := range admins {
	// 		result := database.DB.Create(&admin)
	// 		if result.Error != nil {
	// 			log.Fatalf("Error by seeding admin with ID: %s; Error: %s", admin.ID, result.Error)
	// 		}
	// 	}
	//}

	logger.Log.Info("Migration ran successfully")

	return nil
}
