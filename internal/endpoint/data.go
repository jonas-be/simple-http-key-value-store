package endpoint

import (
	"fmt"
	"net/http"
)

type Database interface {
	Get(key string) string
	Set(key string, value string)
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

	switch method {
	case "GET":
		getData(dh.Db, w, key)
		break
	case "PUT":
		putData(dh.Db, w, key, value)
		break
	case "DELETE":
		deleteData(dh.Db, w, key)
		break
	default:
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
	}
}

func getData(db Database, w http.ResponseWriter, key string) bool {
	if !db.Contains(key) {
		http.Error(w, fmt.Sprintf("no value for key %v", key), http.StatusBadRequest)
		return true
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, db.Get(key))
	return false
}

func putData(db Database, w http.ResponseWriter, key string, value string) bool {
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

func deleteData(db Database, w http.ResponseWriter, key string) bool {
	if !db.Contains(key) {
		http.Error(w, fmt.Sprintf("key %v not exists", key), http.StatusBadRequest)
		return true
	}
	db.Delete(key)
	w.WriteHeader(http.StatusOK)
	return false
}
