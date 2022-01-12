package routers

import (
	"beta_service/handlers"

	"github.com/gorilla/mux"
)

func NewUserRouter(router *mux.Router, userHandler *handlers.UserHandler) (*mux.Router, error) {

	// Create Subrouter for the user model
	userRouter := router.PathPrefix("/redeem").Subrouter()

	userRouter.HandleFunc("/userinfo", userHandler.HandleCreateUser())

	return router, nil
}
