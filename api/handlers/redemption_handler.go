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
func (h *RedemptionHandler) HandleGetRedemptionAssets(ctx *gin.Context) {

	var rr_list []models.RedemptionRequest

	err := ctx.ShouldBindJSON(&rr_list)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	assets, err := h.dbAccess.SelectRedemptionAssets(rr_list)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, assets)
	log.Print(http.StatusOK, nil)
}

// This function parses the KYC form and creates a new KYC entry
func (h *RedemptionHandler) HandlePostKYC(ctx *gin.Context) {

	var kyc models.KYC

	err := ctx.ShouldBindJSON(&kyc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = kyc.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		err = h.dbAccess.InsertKYC(kyc)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {
			ctx.JSON(http.StatusOK, gin.H{"Message": "KYC info successfully added"})
		}
	}
}

// This function parses the Solana redemption program information provided and creates a new Redemption Info entry
func (h *RedemptionHandler) HandlePostRedemptionInfo(ctx *gin.Context) {

	var ri_list []models.RedemptionInfo

	err := ctx.ShouldBindJSON(&ri_list)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	for _, ri := range ri_list {
		err = ri.Validate()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {
			err = h.dbAccess.InsertRedemptionInfo(ri)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
			}
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Redemption info successfully added"})
}
