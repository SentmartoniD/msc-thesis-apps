package handlers

import (
	"fmt"
	"go-food-delivery-app/auth-service/internal/models"
	"go-food-delivery-app/auth-service/internal/services"
	"go-food-delivery-app/auth-service/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type AuthService interface {
	SignUp(newUserData services.CreateUserData) (string, error)
	SignIn(email string, password string) (string, error)
}

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type SignUpRequest struct {
	Email     string          `json:"email" validate:"required"`
	Password  string          `json:"password" validate:"required"`
	UserRole  models.UserRole `json:"user_role" validate:"required"`
	FirstName string          `json:"first_name" validate:"required"`
	LastName  string          `json:"last_name" validate:"required"`
}

func (handler *AuthHandler) SignUpHandler(context *gin.Context) {
	log := logger.Log.Named("SignUpHandler")

	log.Info("started")

	var request SignUpRequest

	err := context.ShouldBind(&request)
	if err != nil {
		log.Warn("error with binding request",
			zap.Error(err),
		)

		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(&request)
	if err != nil {
		log.Warn("request is not valid",
			zap.Error(err),
		)

		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUserData := services.CreateUserData{
		Email:     request.Email,
		Password:  request.Password,
		UserRole:  request.UserRole,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	token, err := handler.authService.SignUp(newUserData)
	if err != nil {
		log.Error("error with signup",
			zap.Error(err),
		)

		if err == fmt.Errorf("email already registered") {
			context.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := SignInResponse{
		Token: token,
	}

	log.Info("finished")

	context.JSON(http.StatusOK, response)
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token string  `json:"token"`
	Error *string `json:"error,omitempty"`
}

func (handler *AuthHandler) SignInHandler(context *gin.Context) {
	log := logger.Log.Named("SignInHandler")

	log.Info("started")

	var request SignInRequest

	err := context.ShouldBind(&request)
	if err != nil {
		log.Warn("error with binding request",
			zap.Error(err),
		)

		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(&request)
	if err != nil {
		log.Warn("request is not valid",
			zap.Error(err),
		)

		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := handler.authService.SignIn(request.Email, request.Password)
	if err != nil {
		log.Error("error with signin",
			zap.Error(err),
		)

		if err == fmt.Errorf("invalid email") || err == fmt.Errorf("invalid password") {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := SignInResponse{
		Token: token,
	}

	log.Info("finished")

	context.JSON(http.StatusOK, response)
}
