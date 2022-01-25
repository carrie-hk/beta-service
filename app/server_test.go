package main

import (
	"net/http"
	"testing"
)

func TestHTTPFails(t *testing.T) {
	_, err := http.Get("http://localhost:443/assets/all")
	if err != nil {
		t.Logf("Successfully failed to send request using HTTP")
		return
	}
}
