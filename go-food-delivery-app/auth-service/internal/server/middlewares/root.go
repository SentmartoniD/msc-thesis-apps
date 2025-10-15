package middlewares

import (
	"fmt"
	"go-food-delivery-app/auth-service/pkg/logger"

	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	environment  = os.Getenv("ENVIRONMENT")
	serverDomain = os.Getenv("SERVER_DOMAIN")
)

func init() {
	environment = strings.ToLower(environment)
	if environment == "" {
		environment = "development"
	}

	if serverDomain == "" {
		serverDomain = "localhost"
	}
}

// middlewareRecovery recovers middleware from a problem
func middlewareRecovery(ctx *gin.Context) {
	log := logger.Log.WithOptions(zap.Fields())

	if err := recover(); err != nil {
		_, file, _, _ := runtime.Caller(2)
		file = filepath.Base(file)
		file = strings.Split(file, ".")[0]

		titleCaser := cases.Title(language.English)
		file = titleCaser.String(file)

		log.Error(fmt.Sprintf("panic recovered in %s Middleware", file),
			zap.String("recover", fmt.Sprintf("%v", err)),
		)

		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}
