package main

import (
	"net/http"
	"simple-http-server/internal/database"
	"simple-http-server/internal/endpoint"
)

var db = endpoint.DataHandler{Db: database.Database{}}

func main() {
	http.HandleFunc("/data", db.HandelRequest)
	http.ListenAndServe(":8080", nil)
}
