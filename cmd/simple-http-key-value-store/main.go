package main

import (
	"net/http"
	"simple-http-key-value-store/internal/database"
	"simple-http-key-value-store/internal/endpoint"
	"simple-http-key-value-store/internal/middelware"
)

var db = endpoint.DataHandler{Db: &database.Database{Data: map[string]string{}}}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/data", db.HandleRequest)
	http.ListenAndServe(":8080", middelware.LoggingMiddleware(mux))
}
