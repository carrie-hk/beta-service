package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
)

var CORS_Middleware = cors.New(cors.Config{
	AllowOrigins:     []string{"https://localhost:3000"},
	AllowMethods:     []string{"GET", "POST"},
	AllowHeaders:     []string{"Content-Type"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	AllowOriginFunc: func(origin string) bool {
		return origin == "https://localhost:3000"
	},
	MaxAge: 12 * time.Hour,
})
