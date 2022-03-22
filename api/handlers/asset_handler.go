package handlers

import (
	"beta_service/api/models"
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

// Return all assets in the database
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

// Return all assets in the database marked as 'featured'
func (h *AssetHandler) HandleGetFeaturedAssets(ctx *gin.Context) {

	assets, err := h.dbAccess.GetFeaturedAssets()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, assets)
	log.Print(http.StatusOK, nil)
}

// Update an individual asset's status
func (h *AssetHandler) HandleUpdateAssetStatus(ctx *gin.Context) {

	var su_list []models.StatusUpdate

	err := ctx.ShouldBindJSON(&su_list)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	for _, su_item := range su_list {
		err := su_item.Validate()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		} else {
			err = h.dbAccess.UpdateAssetStatus(su_item)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succesfully read update info from JSON"})
	log.Print(http.StatusOK, nil)
}
