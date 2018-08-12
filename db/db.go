package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/ericfonseca/watchparty/models"
	_ "github.com/lib/pq"
)

var _db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"

	getVenues   = `SELECT * FROM venues`
	insertVenue = `INSERT INTO venues(address, city, description) VALUES($1, $2, $3)`

	getEvents   = `SELECT * FROM events`
	insertEvent = `INSERT INTO events(title, type, city, start_time) VALUES($1, $2, $3, $4)`

	getUsers   = `SELECT * FROM users`
	insertUser = `INSERT INTO users(name, email) VALUES($1, $2)`

	getWatchers   = `SELECT * FROM watchers`
	insertWatcher = `INSERT INTO watchers(event_id, user_id) VALUES ($1, $2)`

	getHosters   = `SELECT * FROM hosters`
	insertHoster = `INSERT INTO hosters(event_id, venue_id) VALUES ($1, $2)`

	getInterests   = `SELECT * FROM interests`
	insertInterest = `INSERT INTO interests(type, city, user_id) VALUES ($1, $2, $3)`
)

func Init() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	_db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = _db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func Close() {
	_db.Close()
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

func InsertVenue(venue models.Venue) error {
	_, err := _db.Exec(insertVenue, venue.Address, venue.City, venue.Description)
	return err
}

func GetVenues(id, city string) ([]byte, error) {
	q := getVenues
	sep := "WHERE"
	if id != "" {
		q = fmt.Sprintf("%s %s id=%s", q, sep, id)
		sep = "AND"
	}

	if city != "" {
		q = fmt.Sprintf("%s %s city='%s'", q, sep, city)
	}
	log.Printf("query: %s\n", q)
	rows, err := _db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buf := bytes.NewBuffer(nil)
	v := models.Venue{}
	for rows.Next() {
		err := rows.Scan(&v.ID, &v.City, &v.Address, &v.Description)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		b, err := json.Marshal(v)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		buf.Write(b)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func InsertEvent(event models.Event) error {
	_, err := _db.Exec(insertEvent, event.Title, event.Type, event.City, event.StartTime)
	return err
}

func GetEvents(id, city, eventType, startTime string) ([]byte, error) {
	q := getEvents
	sep := "WHERE"
	if id != "" {
		q = fmt.Sprintf("%s %s id=%s", q, sep, id)
		sep = "AND"
	}

	if city != "" {
		q = fmt.Sprintf("%s %s city='%s'", q, sep, city)
	}

	if eventType != "" {
		q = fmt.Sprintf("%s %s type='%s'", q, sep, eventType)
	}

	if startTime != "" {
		q = fmt.Sprintf("%s %s start_time='%s'", q, sep, startTime)
	}
	log.Printf("query: %s\n", q)
	rows, err := _db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buf := bytes.NewBuffer(nil)
	e := models.Event{}
	for rows.Next() {
		err := rows.Scan(&e.ID, &e.Title, &e.Type, &e.City, &e.StartTime)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		b, err := json.Marshal(e)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		buf.Write(b)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func InsertUser(user models.User) error {
	_, err := _db.Exec(insertUser, user.Name, user.Email)
	return err
}

func GetUsers(id string) ([]byte, error) {
	q := getUsers
	sep := "WHERE"
	if id != "" {
		q = fmt.Sprintf("%s %s id=%s", q, sep, id)
		sep = "AND"
	}

	log.Printf("query: %s\n", q)
	rows, err := _db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buf := bytes.NewBuffer(nil)
	u := models.User{}
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		b, err := json.Marshal(u)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		buf.Write(b)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func InsertWatcher(watcher models.Watcher) error {
	_, err := _db.Exec(insertWatcher, watcher.EventID, watcher.UserID)
	return err
}

func GetWatchers(event_id, user_id string) ([]byte, error) {
	q := getWatchers
	sep := "WHERE"
	if event_id != "" {
		q = fmt.Sprintf("%s %s event_id=%s", q, sep, event_id)
		sep = "AND"
	}

	if user_id != "" {
		q = fmt.Sprintf("%s %s user_id='%s'", q, sep, user_id)
	}

	log.Printf("query: %s\n", q)
	rows, err := _db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buf := bytes.NewBuffer(nil)
	w := models.Watcher{}
	for rows.Next() {
		err := rows.Scan(&w.EventID, &w.UserID)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		b, err := json.Marshal(w)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		buf.Write(b)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func InsertHoster(hoster models.Hoster) error {
	_, err := _db.Exec(insertHoster, hoster.EventID, hoster.VenueID)
	return err
}

func GetHosters(event_id, venue_id string) ([]byte, error) {
	q := getHosters
	sep := "WHERE"
	if event_id != "" {
		q = fmt.Sprintf("%s %s event_id=%s", q, sep, event_id)
		sep = "AND"
	}

	if venue_id != "" {
		q = fmt.Sprintf("%s %s venue_id='%s'", q, sep, venue_id)
	}

	log.Printf("query: %s\n", q)
	rows, err := _db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buf := bytes.NewBuffer(nil)
	h := models.Hoster{}
	for rows.Next() {
		err := rows.Scan(&h.EventID, &h.VenueID)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		b, err := json.Marshal(h)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		buf.Write(b)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func InsertInterest(interest models.Interest) error {
	_, err := _db.Exec(insertInterest, interest.Type, interest.City, interest.UserID)
	return err
}

func GetInterests(event_type, city, user_id string) ([]byte, error) {
	q := getInterests
	sep := "WHERE"
	if event_type != "" {
		q = fmt.Sprintf("%s %s type=%s", q, sep, event_type)
		sep = "AND"
	}

	if city != "" {
		q = fmt.Sprintf("%s %s city='%s'", q, sep, city)
	}

	if user_id != "" {
		q = fmt.Sprintf("%s %s user_id='%s'", q, sep, user_id)
	}

	log.Printf("query: %s\n", q)
	rows, err := _db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buf := bytes.NewBuffer(nil)
	i := models.Interest{}
	for rows.Next() {
		err := rows.Scan(i.Type, i.City, i.UserID)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		b, err := json.Marshal(i)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		buf.Write(b)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
