package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCorsMiddleware(allowed_origin string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{allowed_origin},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == allowed_origin
		},
		MaxAge: 12 * time.Hour,
	})
}
