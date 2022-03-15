package router_groups

import (
	"beta_service/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRedemptionRouter(router *gin.RouterGroup, kycHandler *handlers.KycHandler) {

	router.GET("/assets", kycHandler.HandleGetRedemptionInfo)

	// Create router for inserting information into KYC form
	router.POST("/kyc", kycHandler.HandleCreateKYC)
}
