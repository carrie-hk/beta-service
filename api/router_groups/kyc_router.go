package router_groups

import (
	"beta_service/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewUserRouter(router *gin.RouterGroup, userHandler *handlers.KycHandler) {

	// Create router for inserting information into KYC form
	router.POST("/userinfo", userHandler.HandleCreateKYC)
}
