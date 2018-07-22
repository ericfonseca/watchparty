package api

import (
    "fmt"
    "net/http"
)

func VenuesHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        fmt.Printf("YOU POSTY EH?\n")
    } else {
        fmt.Printf("YOU GETTY EH?\n")
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