package router_groups

import (
	"beta_service/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewAssetRouter(router *gin.RouterGroup, assetHandler *handlers.AssetHandler) {

	// Create routers for returning all the assets and a subset of assets
	router.GET("/all", assetHandler.HandleGetAllAssets)
	router.GET("/featured", assetHandler.HandleGetFeaturedAssets)
}
