package handlers

import (
	"beta_service/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssetHandler struct {
	dbAccess *db.DbAccess
}

func NewAssetHandler(dbAccess *db.DbAccess) (*AssetHandler, error) {
	return &AssetHandler{dbAccess: dbAccess}, nil
}

func (h *AssetHandler) HandleTestGetAssets() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Directed to HandleTestGetAssets()")

		assets, err := h.dbAccess.AssetDbAccess.TestAssets()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error getting assets from database")
			return
		}
		b, err := json.Marshal(assets)
		if err != nil {
			log.Println(err)
		}

		fmt.Fprintf(w, "%s", b)
	}

}

//This function returns all of the assets in the AXU.whisky_bottles
func (h *AssetHandler) HandleGetAllAssets(ctx *gin.Context) {
	log.Print("Directed to HandleGetAllAssets()")

	assets, err := h.dbAccess.AssetDbAccess.Assets()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, assets)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Succesfully returned all assets"})
	}

	ctx.JSON(http.StatusOK, nil)
}

//This function returns a featured subset of the bottles
func (h *AssetHandler) HandleGetFeaturedAssets(ctx *gin.Context) {
	log.Print("Directed to HandleGetFeaturedAssets()")

	assets, err := h.dbAccess.AssetDbAccess.FeaturedAssets()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, assets)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Succesfully returned featured assets"})
	}

	ctx.JSON(http.StatusOK, nil)
}
