package endpoint

import (
	"net/http"
	"simple-http-server/internal/database"
	"simple-http-server/internal/endpoint/data"
)

type DataHandler struct {
	Db database.Database
}

func (dataHandler DataHandler) HandelRequest(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	if key == "" {
		http.Error(w, "no key set", http.StatusBadRequest)
		return
	}

	switch method {
	case "GET":
		if data.GetData(dataHandler.Db, w, key) {
			return
		}
		break
	case "PUT":
		if data.PutData(dataHandler.Db, w, key, value) {
			return
		}
		break
	case "DELETE":
		if data.DeleteData(dataHandler.Db, w, key) {
			return
		}
		break
	default:
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
	}
}
