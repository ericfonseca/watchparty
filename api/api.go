package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ericfonseca/watchparty/db"
	"github.com/ericfonseca/watchparty/models"
)

func VenuesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		res, err := db.GetVenues()
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

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	//...
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	//...
}

func watchersHandler(w http.ResponseWriter, r *http.Request) {
	//...
}

func hostersHandler(w http.ResponseWriter, r *http.Request) {
	//...
}
