package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// Set security headers using unrolled secure library
var secureMiddleware = secure.New(secure.Options{
	AllowedHosts:          []string{},
	AllowedHostsAreRegex:  true,
	HostsProxyHeaders:     []string{"X-Forwarded-Host"},
	SSLRedirect:           true,
	SSLHost:               "",
	SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	STSSeconds:            31536000,
	STSIncludeSubdomains:  true,
	STSPreload:            true,
	FrameDeny:             true,
	ContentTypeNosniff:    true,
	BrowserXssFilter:      true,
	ContentSecurityPolicy: "script-src $NONCE",
})

// Convert secureMiddleware options into format usable by Gin
var SecureFunc = func() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			c.Abort()
			return
		}

		// Avoid header rewrite if response is a redirection.
		if status := c.Writer.Status(); status > 300 && status < 399 {
			c.Abort()
		}
	}
}()

var CORS_Middleware = cors.New(cors.Config{
	AllowOrigins:     []string{"https://baxus.co"},
	AllowMethods:     []string{"GET", "POST"},
	AllowHeaders:     []string{"Content-Type"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	AllowOriginFunc: func(origin string) bool {
		return origin == "https://baxus.co"
	},
	MaxAge: 12 * time.Hour,
})
