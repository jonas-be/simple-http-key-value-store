package data

import (
	"fmt"
	"net/http"
	"simple-http-server/internal/database"
)

func DeleteData(db database.Database, w http.ResponseWriter, key string) bool {
	if !db.Contains(key) {
		http.Error(w, fmt.Sprintf("key %v not exists", key), http.StatusBadRequest)
		return true
	}
	db.Delete(key)
	w.WriteHeader(http.StatusOK)
	return false
}
