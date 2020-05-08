package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	srv := &http.Server{
		Addr:         "localhost:8082",
		Handler:      router,
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Millisecond,
		IdleTimeout:  60 * time.Millisecond,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
