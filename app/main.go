package main

import (
	"beta_service/db"
	"beta_service/handlers"
	"beta_service/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
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

	/*** DeBugging Server ***/
	devServer := &http.Server{
		Handler:      router,
		Addr:         ":5050",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	/*** Production Server ***/
	prodServer := &http.Server{
		Handler:      router,
		Addr:         ":5004",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		log.Print("Debugging Server Running and Accepting Requests", devServer.Addr)
		log.Fatal(devServer.ListenAndServe())
		wg.Done()
	}()

	go func() {
		log.Print("Production Server Running and Accepting Requests", prodServer.Addr)
		log.Fatal(prodServer.ListenAndServe())
		wg.Done()
	}()

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

	wg.Wait()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
