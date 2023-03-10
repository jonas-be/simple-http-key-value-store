package endpoint

import (
	"errors"
	"fmt"
	"net/http"
	"simple-http-key-value-store/internal/database"
)

//go:generate mockery --name Database
type Database interface {
	Get(key string) string
	Set(key string, value string) error
	Delete(key string)
	Contains(key string) bool
}

type DataHandler struct {
	Db Database
}

func (dh DataHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	if key == "" {
		http.Error(w, "no key set", http.StatusBadRequest)
		return
	}

	var statusCode int
	var body string
	switch method {
	case "GET":
		statusCode, body = getData(dh.Db, key)
		break
	case "PUT":
		statusCode, body = putData(dh.Db, key, value)
		break
	case "DELETE":
		statusCode, body = deleteData(dh.Db, key)
		break
	default:
		statusCode, body = http.StatusMethodNotAllowed, "unsupported method"
	}
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "%s\n", body)
}

func getData(db Database, key string) (int, string) {
	if !db.Contains(key) {
		return http.StatusBadRequest, fmt.Sprintf("no value for key %v", key)
	}
	return http.StatusOK, db.Get(key)
}

func putData(db Database, key string, value string) (int, string) {
	if value == "" {
		return http.StatusBadRequest, "no value set"
	}
	isNewEntry := !db.Contains(key)
	err := db.Set(key, value)
	if err != nil {
		if errors.As(err, &database.OutOfStorageError{}) {
			return http.StatusInsufficientStorage, err.Error()
		} else if errors.As(err, &database.InputError{}) {
			return http.StatusRequestEntityTooLarge, err.Error()
		} else {
			return http.StatusInternalServerError, err.Error()
		}
	}
	if isNewEntry {
		return http.StatusCreated, ""
	}
	return http.StatusOK, ""
}

func deleteData(db Database, key string) (int, string) {
	if !db.Contains(key) {
		return http.StatusBadRequest, fmt.Sprintf("key %v not exists", key)
	}
	db.Delete(key)
	return http.StatusNoContent, ""
}
