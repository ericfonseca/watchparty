package main

import (
	"log"
	"net/http"

	"github.com/ericfonseca/watchparty/api"
	"github.com/ericfonseca/watchparty/db"
)

func main() {
	db.Init()
	defer db.Close()
	http.HandleFunc("/api/venues", api.VenuesHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
