package routers

import (
	"beta_service/handlers"

	"github.com/gin-gonic/gin"
)

func NewUserRouter(router *gin.RouterGroup, userHandler *handlers.UserHandler) {

	// Create router for inserting information into KYC form
	router.GET("/userinfo", userHandler.HandleCreateUser)
}
