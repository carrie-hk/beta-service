package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCorsMiddleware(allowed_origins []string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     allowed_origins,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			for _, allowed_origin := range allowed_origins {
				if allowed_origin == origin {
					return true
				}
			}
			return false
		},
		MaxAge: 12 * time.Hour,
	})
}
