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
	r.HandleFunc("/api/venues/{id}", api.VenueByIDHandler)
	r.HandleFunc("/api/events", api.EventsHandler)
	r.HandleFunc("/api/events/{id}", api.EventByIDHandler)
	r.HandleFunc("/api/users", api.UsersHandler)
	r.HandleFunc("/api/users/{id}", api.UserByIDHandler)
	r.HandleFunc("/api/watchers", api.WatchersHandler)
	r.HandleFunc("/api/hosters", api.HostersHandler)
	r.HandleFunc("/api/interests", api.InterestsHandler)
	http.ListenAndServe(":8000", r)
}
