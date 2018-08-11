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

	getVenues       = `SELECT * FROM venues`
	getVenuesByCity = `SELECT * FROM venues WHERE city=$1`
	getVenueByID    = `SELECT * FROM venues WHERE id=$1`
	insertVenue     = `INSERT INTO venues(address, city, description) VALUES($1, $2, $3)`
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

func GetVenues() ([]byte, error) {
	rows, err := _db.Query(getVenues)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buf := bytes.NewBuffer(nil)
	v := models.Venue{}
	var id int
	for rows.Next() {
		err := rows.Scan(&id, &v.City, &v.Address, &v.Description)
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

func GetVenuesByCity(city string) ([]byte, error) {
	rows, err := _db.Query(getVenuesByCity, city)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buf := bytes.NewBuffer(nil)
	v := models.Venue{}
	var id int
	for rows.Next() {
		err := rows.Scan(&id, &v.City, &v.Address, &v.Description)
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

func GetVenueByID(venueID string) ([]byte, error) {
	rows, err := _db.Query(getVenueByID, venueID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buf := bytes.NewBuffer(nil)
	v := models.Venue{}
	var id int
	for rows.Next() {
		err := rows.Scan(&id, &v.City, &v.Address, &v.Description)
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
