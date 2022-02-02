package middlewares

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
)

var CORS_Middleware = cors.New(cors.Config{
	AllowOrigins:     []string{os.Getenv("ALLOWED_ORIGIN_1")},
	AllowMethods:     []string{"GET", "POST"},
	AllowHeaders:     []string{"Content-Type"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	AllowOriginFunc: func(origin string) bool {
		return origin == os.Getenv("ALLOWED_ORIGIN_1")
	},
	MaxAge: 12 * time.Hour,
})
