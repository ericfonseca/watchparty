package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ericfonseca/watchparty/db"
	"github.com/ericfonseca/watchparty/models"
	"github.com/gorilla/mux"
)

func VenuesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		city := r.URL.Query().Get("city")
		res, err := db.GetVenues("", city)

		if err != nil {
			log.Print("could not get venues", "err", err.Error())
			w.WriteHeader(500)
			return
		}
		w.Write(res)

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("could not read req body", "err", err.Error())
			w.WriteHeader(404)
			return
		}
		venue := models.Venue{}
		err = json.Unmarshal(body, &venue)
		if err != nil {
			w.WriteHeader(404)
			log.Print("could not unmarshall", "err", err.Error())
			return
		}
		err = db.InsertVenue(venue)
		if err != nil {
			w.WriteHeader(404)
			log.Print("could not insert venue", "err", err.Error())
			return
		}
		w.WriteHeader(200)
	}
}

func VenueByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(404)
		log.Print("could not get id from path")
	}

	res, err := db.GetVenues(id, "")
	if err != nil {
		log.Print("could not get venues", "err", err.Error())
		w.WriteHeader(500)
		return
	}
	w.Write(res)
}

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		city := r.URL.Query().Get("city")
		eventType := r.URL.Query().Get("type")
		startTime := r.URL.Query().Get("start")

		res, err := db.GetEvents("", city, eventType, startTime)

		if err != nil {
			log.Print("could not get events", "err", err.Error())
			w.WriteHeader(500)
			return
		}
		w.Write(res)

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("could not read req body", "err", err.Error())
			w.WriteHeader(404)
			return
		}
		event := models.Event{}
		err = json.Unmarshal(body, &event)
		if err != nil {
			w.WriteHeader(404)
			log.Print("could not unmarshall", "err", err.Error())
			return
		}
		err = db.InsertEvent(event)
		if err != nil {
			w.WriteHeader(404)
			log.Print("could not insert event", "err", err.Error())
			return
		}
		w.WriteHeader(200)
	}
}

func EventByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(404)
		log.Print("could not get id from path")
	}

	res, err := db.GetEvents(id, "", "", "")
	if err != nil {
		log.Print("could not get events", "err", err.Error())
		w.WriteHeader(500)
		return
	}
	w.Write(res)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		res, err := db.GetUsers("")

		if err != nil {
			log.Print("could not get users", "err", err.Error())
			w.WriteHeader(500)
			return
		}
		w.Write(res)

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("could not read req body", "err", err.Error())
			w.WriteHeader(404)
			return
		}
		user := models.User{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			w.WriteHeader(404)
			log.Print("could not unmarshall", "err", err.Error())
			return
		}
		err = db.InsertUser(user)
		if err != nil {
			w.WriteHeader(404)
			log.Print("could not insert event", "err", err.Error())
			return
		}
		w.WriteHeader(200)
	}
}

func UserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(404)
		log.Print("could not get id from path")
	}

	res, err := db.GetUsers(id)
	if err != nil {
		log.Print("could not get users", "err", err.Error())
		w.WriteHeader(500)
		return
	}
	w.Write(res)
}

func WatchersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		eventID := r.URL.Query().Get("event_id")
		userID := r.URL.Query().Get("user_id")

		res, err := db.GetWatchers(eventID, userID)

		if err != nil {
			log.Print("could not get watchers", "err", err.Error())
			w.WriteHeader(500)
			return
		}
		w.Write(res)

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("could not read req body", "err", err.Error())
			w.WriteHeader(404)
			return
		}
		watcher := models.Watcher{}
		err = json.Unmarshal(body, &watcher)
		if err != nil {
			w.WriteHeader(404)
			log.Print("could not unmarshall", "err", err.Error())
			return
		}
		err = db.InsertWatcher(watcher)
		if err != nil {
			w.WriteHeader(404)
			log.Print("could not insert watcher", "err", err.Error())
			return
		}
		w.WriteHeader(200)
	}
}

func HostersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		eventID := r.URL.Query().Get("event_id")
		venueID := r.URL.Query().Get("venue_id")

		res, err := db.GetHosters(eventID, venueID)

		if err != nil {
			log.Print("could not get hosters", "err", err.Error())
			w.WriteHeader(500)
			return
		}
		w.Write(res)

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("could not read req body", "err", err.Error())
			w.WriteHeader(404)
			return
		}
		hoster := models.Hoster{}
		err = json.Unmarshal(body, &hoster)
		if err != nil {
			w.WriteHeader(404)
			log.Print("could not unmarshall", "err", err.Error())
			return
		}
		err = db.InsertHoster(hoster)
		if err != nil {
			w.WriteHeader(404)
			log.Print("could not insert hoster", "err", err.Error())
			return
		}
		w.WriteHeader(200)
	}
}

func InterestsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		eventType := r.URL.Query().Get("type")
		city := r.URL.Query().Get("city")
		userID := r.URL.Query().Get("user_id")

		res, err := db.GetInterests(eventType, city, userID)

		if err != nil {
			log.Print("could not get interests", "err", err.Error())
			w.WriteHeader(500)
			return
		}
		w.Write(res)

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("could not read req body", "err", err.Error())
			w.WriteHeader(404)
			return
		}
		interest := models.Interest{}
		err = json.Unmarshal(body, &interest)
		if err != nil {
			w.WriteHeader(404)
			log.Print("could not unmarshall", "err", err.Error())
			return
		}
		err = db.InsertInterest(interest)
		if err != nil {
			w.WriteHeader(404)
			log.Print("could not insert interest", "err", err.Error())
			return
		}
		w.WriteHeader(200)
	}
}
