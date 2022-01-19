package routers

import (
	"beta_service/handlers"

	"github.com/gin-gonic/gin"
)

func NewAssetRouter(router *gin.RouterGroup, assetHandler *handlers.AssetHandler) {

	// Create routers for returning all the assets and a subset of assets
	router.GET("/all", assetHandler.HandleGetAllAssets)
	router.GET("/featured", assetHandler.HandleGetFeaturedAssets)
}
