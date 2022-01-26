package main

import (
	"beta_service/db"
	"beta_service/handlers"
	"beta_service/routers"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize database connection and model stores
	dbAccess, err := db.NewDbAccess()
	logFatal(err)

	// Initialize handlers for models
	assetHandler, err := handlers.NewAssetHandler(dbAccess)
	logFatal(err)

	userHandler, err := handlers.NewUserHandler(dbAccess)
	logFatal(err)

	// Create a router
	router := gin.Default()

	// Initialize router groups for handlers
	routers.NewAssetRouter(router.Group("/assets"), assetHandler)
	routers.NewUserRouter(router.Group("/redeem"), userHandler)

	tls_cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, // Need this for HTTP/2
		},
	}

	/*** DeBugging Server ***/
	devServer := &http.Server{
		Handler:      router,
		Addr:         ":5050",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		TLSConfig:    tls_cfg,
	}

	/*** Production Server ***/
	prodServer := &http.Server{
		Handler:      router,
		Addr:         ":443",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		TLSConfig:    tls_cfg,
	}

	wg := new(sync.WaitGroup)
	wg.Add(3)

	// For production, we will need to use a more appropriate certificate-keypair combination obtained from Let's Encrypt through Github
	// For now, since we're using a self-signed certificate, we must use curl with the -k flag in order to complete the request
	go func() {
		log.Print("Debugging Server Running and Accepting Requests", devServer.Addr)
		log.Fatal(devServer.ListenAndServeTLS("../server.rsa.crt", "../server.rsa.key"))
		wg.Done()
	}()

	go func() {
		log.Print("Production Server Running and Accepting Requests", prodServer.Addr)
		log.Fatal(prodServer.ListenAndServeTLS("../server.rsa.crt", "../server.rsa.key"))
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
