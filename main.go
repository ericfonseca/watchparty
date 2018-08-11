package main

import (
	"net/http"

	"github.com/ericfonseca/watchparty/api"
	"github.com/ericfonseca/watchparty/db"
	"github.com/gorilla/mux"
)

func main() {
	db.Init()
	defer db.Close()
	r := mux.NewRouter()
	r.HandleFunc("/api/venues", api.VenuesHandler)
	r.HandleFunc("/api/venues/{id}", api.VenuesByIDHandler)
	http.ListenAndServe(":8181", r)
}
