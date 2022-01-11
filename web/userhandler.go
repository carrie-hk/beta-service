package web

import (
	"beta_service/db"
	"beta_service/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type UserHandler struct {
	router *mux.Router
	store  *db.Store
}

func NewUserHandler(store *db.Store, router *mux.Router) *UserHandler {
	h := &UserHandler{
		router: router,
		store:  store,
	}

	subrouter := router.PathPrefix("/redeem").Subrouter()
	subrouter.HandleFunc("/kycform", h.HandleCreateUser())

	return h
}

//This function parses the KYC form and creates a new user
func (h *UserHandler) HandleCreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Directed to Handle Create User")

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(w, "ParseForm() err: %v", err)
		}
		decoder := schema.NewDecoder()
		var user models.User
		err = decoder.Decode(user, r.PostForm)
		if err != nil {
			log.Println(w, "DecodeForm err: %v", err)
		}
		err = h.store.UserStore.CreateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error getting assets from database")
			return
		}

		fmt.Fprintf(w, "%s", "Success!")

	}

}
