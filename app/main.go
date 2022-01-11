package main

import (
	"beta_service/db"
	"beta_service/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	// Open database connection
	database, err := db.Open()
	log.Print("Database Opened")
	if err != nil {
		print(err)
	}

	defer database.Close()

	h := handlers.NewHandler(database)

	r := mux.NewRouter()
	r.HandleFunc("/assets", h.HandleGetAssets).Methods("GET")
	r.HandleFunc("/", h.HandleGetFeaturedAssets).Methods("GET")
	r.HandleFunc("/kycform", h.HandleCreateUser).Methods("POST")

	s := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:5000",
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

	log.Print("Server Running and Accepting Requests")
	log.Fatal(s.ListenAndServe())

}
