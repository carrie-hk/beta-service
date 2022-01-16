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

	s := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:5000",
		TLSConfig:    tls_cfg,
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
	// For production, we will need to use a more appropriate certificate-keypair combination obtained from Let's Encrypt through Github
	// For now, since we're using a self-signed certificate, we must use curl with the -k flag in order to complete the request
	log.Fatal(s.ListenAndServeTLS("server.rsa.crt", "server.rsa.key"))

}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
