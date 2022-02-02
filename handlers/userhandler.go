package handlers

import (
	"beta_service/db"
	"beta_service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	dbAccess *db.DbAccess
}

func NewUserHandler(dbAccess *db.DbAccess) (*UserHandler, error) {
	return &UserHandler{dbAccess: dbAccess}, nil
}

//This function parses the KYC form and creates a new user
func (h *UserHandler) HandleCreateUser(ctx *gin.Context) {

	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Succesfully read user info from JSON"})
	}

	err = h.dbAccess.UserDbAccess.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Succesfully added user to database"})
	}

	ctx.JSON(http.StatusOK, nil)
}
