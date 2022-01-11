package main

import (
	"beta_service/db"
	"beta_service/web"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	//Initialize database connection and model stores
	store, err := db.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	//Initialize router/mutex/handler for models
	router, err := web.NewHandler(store)
	if err != nil {
		log.Fatal(err)
	}

	s := &http.Server{
		Handler:      router.Router,
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
