package router_groups

import (
	"beta_service/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRedemptionRouter(router *gin.RouterGroup, redemptionHandler *handlers.RedemptionHandler) {

	// This is not traditional/idiomatic REST practice - typically, you'd use a GET request for this purpose, but we are using a POST request for a few reasons:
	//
	// 1) We want to be able to easily return multiple assets with one single HTTP request instead of using an extremely lengthy URL with multiple query parameters (especially since
	// each parameter is a 32 character public key)
	// 2) We don't want to put non-trivial data, such as an asset's mint address, in the URL
	//
	router.POST("/assets", redemptionHandler.HandleGetRedemptionInfo)

	// Create router for inserting information into KYC form
	router.POST("/kyc", redemptionHandler.HandleCreateKYC)
}
