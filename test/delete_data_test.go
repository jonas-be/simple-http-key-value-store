package test

import (
	"net/http"
	"net/http/httptest"
	"simple-http-key-value-store/internal/endpoint"
	"simple-http-key-value-store/test/util"
	"testing"
)

func TestDeleteOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/data?key=c", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: util.MockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	util.AssertStatusCode(t, resp, http.StatusOK)
	util.AssertBody(t, resp, "")

	if dataHandler.Db.Contains("c") {
		t.Error("expected \"c\" not exists, got \"c\" exists")
	}
}

func TestDeleteBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/data?key=z", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: util.MockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	util.AssertStatusCode(t, resp, http.StatusBadRequest)
	util.AssertBody(t, resp, "key z not exists\n")
}
