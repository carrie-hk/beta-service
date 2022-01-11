package main

import (
	"beta_service/db"
	"beta_service/web"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAssetsHandler(t *testing.T) {

	/******* STRING UP SERVER *******/

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

	log.Print("Server Running and Accepting Requests")
	log.Fatal(s.ListenAndServe())

	tt := []struct {
		name       string
		method     string
		path       string
		want       string
		statusCode int
	}{
		{
			name:       "return all assets",
			method:     http.MethodGet,
			path:       `http://127.0.0.1:5000\assets`,
			want:       `[{"id":1,"name":"Foo","price":10}]`,
			statusCode: http.StatusOK,
		},
		{
			name:       "return featured assets",
			method:     http.MethodPost,
			path:       `http://127.0.0.1:5000\`,
			want:       "Method not allowed",
			statusCode: http.StatusOK,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			request, err := http.NewRequest(tc.method, tc.path, nil)
			responseRecorder := httptest.NewRecorder()

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
