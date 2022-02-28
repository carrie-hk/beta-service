package main

import (
	"beta_service/db_access"
	"beta_service/handlers"
	"beta_service/middlewares"
	"beta_service/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Enable use of .env file if running with "local" flag
	// parseBuildArgs()
	godotenv.Load(".env")

	// Initialize database connection and model stores
	dbAccess, err := db_access.NewDbAccess()
	logFatal(err)

	// Initialize handlers for models
	assetHandler, err := handlers.NewAssetHandler(dbAccess)
	logFatal(err)

	userHandler, err := handlers.NewUserHandler(dbAccess)
	logFatal(err)

	// Create a router
	router := gin.Default()

	router.Use(
		// Using the hard-coded CORS function in middlewares instead of the CORS object provided by gin-contrib/cors
		middlewares.CORS_Middleware(os.Getenv("BAXUS_ORIGIN")),
	)

	// Initialize router groups for handlers
	routers.NewAssetRouter(router.Group("/assets"), assetHandler)
	routers.NewUserRouter(router.Group("/redeem"), userHandler)

	server := &http.Server{
		Handler:      router,
		Addr:         os.Getenv("SERVER_PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		log.Print("Server Running and Accepting Requests", server.Addr)
		log.Fatal(server.ListenAndServe())
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

	log.Print("Server Running and Accepting Requests")

	wg.Wait()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
