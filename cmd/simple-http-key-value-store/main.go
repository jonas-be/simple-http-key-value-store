package main

import (
	"net/http"
	"simple-http-key-value-store/internal/database"
	"simple-http-key-value-store/internal/endpoint"
)

var db = endpoint.DataHandler{Db: &database.Database{}}

func main() {
	http.HandleFunc("/data", db.HandleRequest)
	http.ListenAndServe(":8080", nil)
}
