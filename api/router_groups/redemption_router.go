package router_groups

import (
	"beta_service/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRedemptionRouter(router *gin.RouterGroup, redemptionHandler *handlers.RedemptionHandler) {

	// This is not traditional/idiomatic REST practice - typically, you'd use a GET request for this purpose, but we are using a POST request for a few reasons:
	// 1)
	// 2)
	router.POST("/assets", redemptionHandler.HandleGetRedemptionInfo)

	// Create router for inserting information into KYC form
	router.POST("/kyc", redemptionHandler.HandleCreateKYC)
}
