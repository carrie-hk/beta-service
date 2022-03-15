package router_groups

import (
	"beta_service/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRedemptionRouter(router *gin.RouterGroup, redemptionHandler *handlers.RedemptionHandler) {

	router.GET("/assets", redemptionHandler.HandleGetRedemptionInfo)

	// Create router for inserting information into KYC form
	router.POST("/kyc", redemptionHandler.HandleCreateKYC)
}
