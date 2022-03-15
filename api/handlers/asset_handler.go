package handlers

import (
	"beta_service/db/data_access"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AssetHandler struct {
	dbAccess *data_access.DbAccess
}

func NewAssetHandler(dbAccess *data_access.DbAccess) (*AssetHandler, error) {
	return &AssetHandler{dbAccess: dbAccess}, nil
}

//This function returns all of the assets in the AXU.whisky_bottles
func (h *AssetHandler) HandleGetAllAssets(ctx *gin.Context) {

	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	pageIndex, err := strconv.Atoi(ctx.Query("pageIdx"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	assets, err := h.dbAccess.GetAllAssets(pageIndex, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, assets)

	log.Print(http.StatusOK, nil)
}

//This function returns a featured subset of the bottles
func (h *AssetHandler) HandleGetFeaturedAssets(ctx *gin.Context) {

	assets, err := h.dbAccess.GetFeaturedAssets()
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
