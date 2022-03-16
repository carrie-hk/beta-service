package handlers

import (
	"beta_service/api/models"
	"beta_service/db/data_access"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RedemptionHandler struct {
	dbAccess *data_access.DbAccess
}

func NewRedemptionHandler(dbAccess *data_access.DbAccess) (*RedemptionHandler, error) {
	return &RedemptionHandler{dbAccess: dbAccess}, nil
}

// This function returns a set of information about redemption for the connected wallet's AXUs
func (h *RedemptionHandler) HandleGetRedemptionInfo(ctx *gin.Context) {

	mintAddr := ctx.Query("mint")
	assets, err := h.dbAccess.GetRedemptionAssets(mintAddr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, assets)

	log.Print(http.StatusOK, nil)
}

// This function parses the KYC form and creates a new KYC entry
func (h *RedemptionHandler) HandleCreateKYC(ctx *gin.Context) {

	var kyc models.KYC

	err := ctx.ShouldBindJSON(&kyc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = kyc.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.dbAccess.CreateKYC(kyc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusOK, nil)
}
