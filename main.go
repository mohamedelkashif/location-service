package main

import (
	"github.com/gorilla/mux"
	"github.com/mohamedelkashif/store-location-service/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/stores").HandlerFunc(handler.GetStores)
	sub.Methods("POST").Path("/stores").HandlerFunc(handler.SaveStore)
	sub.Methods("GET").Path("/stores/").Queries("country", "{country}", "max", "{max}").HandlerFunc(handler.GetStoresByCountry)

	log.Fatal(http.ListenAndServe(":8080", router))
}