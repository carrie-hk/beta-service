package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CORS_Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// I'm pushing this utility function to the branch so that, should we decide it's simpler, we can concatenate the allowed origin strings from env.list,
// then change CORS_Middleware() to CORS_Middleware(allowed_origins string) and .Set("Access-Control-Allow-Origin", allowed_origins),
// then in main call middlewares.concatAllowedOrigins(os.GetEnv("BAXUS_ORIGIN_1"), os.GetEnv...) to concatenate the origins together
// Thoughts?

func concatAllowedOrigins(allowed_origins []string) string {
	var concat_string = ""
	for _, origin_string := range allowed_origins {
		concat_string = concat_string + origin_string + ", "
	}
	return concat_string
}
