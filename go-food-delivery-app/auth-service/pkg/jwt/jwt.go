package jwt

import (
	"errors"
	"go-food-delivery-app/auth-service/pkg/logger"
	"log"
	"os"
	"strconv"
	"strings"

	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type JWTConfig struct {
	// Temporary is a token before the authenticator app is verified
	TemporaryTokenExpiration time.Duration
	JWTTemporarySecret       string

	// Token is the session token when authenticator app is validated
	TokenExpiration time.Duration
	JWTSecret       string
}

var JWT JWTConfig

func LoadJWTConfig() {
	JWT.JWTSecret = os.Getenv("JWT_SECRET")
	if JWT.JWTSecret == "" {
		log.Fatal("auth secret cannot be empty")
	}

	tokenExpirationMinutes, _ := strconv.Atoi(os.Getenv("TOKEN_EXP_MINS"))
	if tokenExpirationMinutes == 0 {
		tokenExpirationMinutes = 48 * 60
	}
	JWT.TokenExpiration = time.Minute * time.Duration(tokenExpirationMinutes)

	JWT.JWTTemporarySecret = os.Getenv("JWT_TEMP_SECRET")
	if JWT.JWTTemporarySecret == "" {
		log.Fatal("temporary auth secret cannot be empty")
	}

	temporaryTokenExpirationMinutes, _ := strconv.Atoi(os.Getenv("TEMP_TOKEN_EXP_MINS"))
	if temporaryTokenExpirationMinutes == 0 {
		temporaryTokenExpirationMinutes = 10
	}
	JWT.TemporaryTokenExpiration = time.Minute * time.Duration(temporaryTokenExpirationMinutes)
}

// GenerateTemporaryToken generates temporary token, it's a token before authenticator code is verified
func GenerateTemporaryToken(userID string, email string, role string) (token string, err error) {
	return TokenGenerate(userID, email, JWT.JWTTemporarySecret, JWT.TemporaryTokenExpiration, role)
}

// ValidateTemporaryToken validates temporary token
func ValidateTemporaryToken(tokenVal string) (*CustomClaims, error) {
	return TokenValidate(tokenVal, JWT.JWTTemporarySecret)
}

// GenerateSessionToken this token is generated when authenticator code is also verified and it is used for API access
func GenerateSessionToken(userID string, email string, role string) (token string, err error) {
	return TokenGenerate(userID, email, JWT.JWTSecret, JWT.TokenExpiration, role)
}

// ValidateSessionToken validates session token
func ValidateSessionToken(tokenVal string) (*CustomClaims, error) {
	return TokenValidate(tokenVal, JWT.JWTSecret)
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	Role  string `json:"role"`
}

// TokenGenerate generates token with a given secret
func TokenGenerate(userID string, email, jwtSecret string, tokenExpiration time.Duration, role string) (token string, err error) {
	tokenUUID := uuid.New().String()

	// Generate a session token
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiration)),
			ID:        tokenUUID,
		},
		Email: email,
		Role:  role,
	}).
		SignedString([]byte(jwtSecret)) // Assuming you have a separate key for session tokens
	if err != nil {
		logger.Log.Error("Failed to generate token", zap.Error(err))
		return
	}

	return
}

func TokenValidate(tokenVal string, jwtSecret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenVal, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method in token")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GetTokenFromEmail(bearer string) (email string, err error) {

	jwtToken := strings.Split(bearer, " ")[1]

	var claims *CustomClaims

	claims, err = ValidateSessionToken(jwtToken)
	if err != nil {
		return
	}

	email = claims.Email

	return
}

func GetTemporaryTokenFromEmail(bearer string) (email string, err error) {

	jwtToken := strings.Split(bearer, " ")[1]

	var claims *CustomClaims

	claims, err = ValidateTemporaryToken(jwtToken)
	if err != nil {
		return
	}

	email = claims.Email

	return
}
