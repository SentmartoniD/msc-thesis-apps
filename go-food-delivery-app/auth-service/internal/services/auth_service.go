package services

import (
	"errors"
	"fmt"
	"go-food-delivery-app/auth-service/internal/models"
	"go-food-delivery-app/auth-service/pkg/crypto"
	"go-food-delivery-app/auth-service/pkg/jwt"
	"go-food-delivery-app/auth-service/platform/database"

	"gorm.io/gorm"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

type CreateUserData struct {
	Email     string
	Password  string
	UserRole  models.UserRole
	FirstName string
	LastName  string
}

func (service *AuthService) SignUp(newUserData CreateUserData) (string, error) {

	// var count int64
	// database.DB.Model(&models.UserCredentials{}).Where("email = ?", newUserData.Email).Count(&count)

	// if count > 0 {
	// 	return "", fmt.Errorf("email already registered")
	// }

	//hashedPassword := crypto.HashPassword(password)

	// TODO:
	// CALL THE USER SERVICE FOR CREATING A NEW USER

	// token, err := jwt.GenerateSessionToken(userCredentials.ID, userCredentials.Email, string(models.Admin))
	// if err != nil {
	// 	return "", fmt.Errorf("failed to generate JWT token: %w", err)
	// }

	return "token", nil
}

func (service *AuthService) SignIn(email string, password string) (string, error) {
	var userCredentials *models.UserCredentials
	result := database.DB.First(&userCredentials, "email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("invalid email")
		}
		return "", result.Error
	}

	valid := crypto.IsPasswordValid(userCredentials.Password, password)
	if !valid {
		return "", fmt.Errorf("invalid password")
	}

	token, err := jwt.GenerateSessionToken(userCredentials.ID, userCredentials.Email, string(models.Admin))
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %w", err)
	}

	return token, nil
}
