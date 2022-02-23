package handlers

import (
	"beta_service/db_access"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssetHandler struct {
	dbAccess *db_access.DbAccess
}

func NewAssetHandler(dbAccess *db_access.DbAccess) (*AssetHandler, error) {
	return &AssetHandler{dbAccess: dbAccess}, nil
}

//This function returns all of the assets in the AXU.whisky_bottles
func (h *AssetHandler) HandleGetAllAssets(ctx *gin.Context) {

	assets, err := h.dbAccess.AssetDbAccess.Assets()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, assets)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		log.Print("Message successful")
	}

	log.Print(http.StatusOK, nil)
}

//This function returns a featured subset of the bottles
func (h *AssetHandler) HandleGetFeaturedAssets(ctx *gin.Context) {

	assets, err := h.dbAccess.AssetDbAccess.FeaturedAssets()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, assets)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		log.Print("Message successful")
	}

	log.Print(http.StatusOK, nil)
}
