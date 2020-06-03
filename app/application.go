package app

import (
	"github.com/gorilla/mux"
	"github.com/nishant01/mybookstore_items-api/clients/elasticsearch"
	"net/http"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
