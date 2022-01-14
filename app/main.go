package main

import (
	"beta_service/db"
	"beta_service/handlers"
	"beta_service/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	// Initialize database connection and model stores
	dbAccess, err := db.NewDbAccess()
	logFatal(err)

	// Create a router
	router := mux.NewRouter()

	// Initialize handlers for models
	assetHandler, err := handlers.NewAssetHandler(dbAccess)
	logFatal(err)

	userHandler, err := handlers.NewUserHandler(dbAccess)
	logFatal(err)

	// Initialize subrouters for handlers
	router, err = routers.NewAssetRouter(router, assetHandler)
	logFatal(err)

	router, err = routers.NewUserRouter(router, userHandler)
	logFatal(err)

	s := &http.Server{
		Handler:      router,
		Addr:         ":5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Print("Server Open")

	// Listen for interrupt or terminal signal from the OS (e.g. Command+C)
	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-shutdown
		switch sig {
		case os.Interrupt:
			log.Print("Server interrupted")
			os.Exit(0)
		case syscall.SIGINT:
			log.Print("Server Cancelled")
			os.Exit(0)
		}
	}()

	log.Print("Server Running and Accepting Requests", s.Addr)
	log.Fatal(s.ListenAndServe())

}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
