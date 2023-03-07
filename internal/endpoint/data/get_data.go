package data

import (
	"fmt"
	"net/http"
	"simple-http-server/internal/database"
)

func GetData(db database.Database, w http.ResponseWriter, key string) bool {
	if !db.Contains(key) {
		http.Error(w, fmt.Sprintf("no value for key %v", key), http.StatusBadRequest)
		return true
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, db.Get(key))
	return false
}
