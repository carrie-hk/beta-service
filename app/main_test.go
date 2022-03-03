package main

import (
	"beta_service/db_access"
	"beta_service/handlers"
	"beta_service/routers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Test_GetAllAssets(t *testing.T) {
	godotenv.Load("../.env")

	dbAccess, err := db_access.NewDbAccess()
	if err != nil {
		t.Fatal(err)
	}

	assetHandler, err := handlers.NewAssetHandler(dbAccess)
	if err != nil {
		t.Fatal(err)
	}

	// Create test router
	test_router := gin.New()
	// Add /asset endpoints to test router
	routers.NewAssetRouter(test_router.Group("/assets"), assetHandler)

	// Create a new server using the test router
	ts := httptest.NewServer(test_router)
	defer ts.Close()

	// Ping the "/assets/all" endpoint
	res, err := http.Get(ts.URL + "/assets/all?pageSize=4&pageIdx=4")
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

func Test_GetFeaturedAssets(t *testing.T) {
	err := godotenv.Load("../.env")

	dbAccess, err := db_access.NewDbAccess()
	if err != nil {
		t.Fatal(err)
	}

	assetHandler, err := handlers.NewAssetHandler(dbAccess)
	if err != nil {
		t.Fatal(err)
	}

	// Create test router
	test_router := gin.New()
	// Add /asset endpoints to test router
	routers.NewAssetRouter(test_router.Group("/assets"), assetHandler)

	// Create a new server using the test router
	ts := httptest.NewServer(test_router)
	defer ts.Close()

	// Ping the "/assets/all" endpoint
	res, err := http.Get(ts.URL + "/assets/featured")
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}