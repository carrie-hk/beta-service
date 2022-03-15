package handlers

import (
	"beta_service/api/models"
	"beta_service/db/data_access"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KycHandler struct {
	dbAccess *data_access.DbAccess
}

func NewKycHandler(dbAccess *data_access.DbAccess) (*KycHandler, error) {
	return &KycHandler{dbAccess: dbAccess}, nil
}

//This function parses the KYC form and creates a new user
func (h *KycHandler) HandleCreateKYC(ctx *gin.Context) {

	var kyc models.KYC

	err := ctx.ShouldBindJSON(&kyc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Succesfully read user info from JSON"})
	}

	err = h.dbAccess.CreateKYC(kyc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Succesfully added KYC information to database"})
	}

	ctx.JSON(http.StatusOK, nil)
}
