package data

import (
	"net/http"
	"simple-http-key-value-store/internal/database"
)

func PutData(db database.Database, w http.ResponseWriter, key string, value string) bool {
	if value == "" {
		http.Error(w, "no value set", http.StatusBadRequest)
		return true
	}
	if db.Contains(key) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	db.Set(key, value)
	return false
}
