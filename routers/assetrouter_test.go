package routers

import (
	"net/http"
	"testing"
)

// Verify that the server successfully returns a GET request from the assets/all URL
func TestGetAllAssets(t *testing.T) {
	_, err := http.Get("http://127.0.0.1:5000/assets/all")
	if err != nil {
		t.Fatal(err)
	}
}

// Verify that the server successfully returns a GET request from the assets/featured URL
func TestGetFeaturedAssets(t *testing.T) {
	_, err := http.Get("http://127.0.0.1:5000/assets/featured")
	if err != nil {
		t.Fatal(err)
	}
}
