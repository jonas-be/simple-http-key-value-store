package test

import (
	"net/http"
	"net/http/httptest"
	"simple-http-server/internal/endpoint"
	"testing"
)

func TestPutOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/data?key=b&value=ZZZ", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: mockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	assertStatusCode(t, resp, http.StatusOK)
	assertBody(t, resp, "")

	assertDBEntry(t, dataHandler.Db, "b", "ZZZ")
}

func TestPutCreated(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/data?key=x&value=X", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: mockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	assertStatusCode(t, resp, http.StatusCreated)
	assertBody(t, resp, "")

	assertDBEntry(t, dataHandler.Db, "x", "X")
}

func TestPutBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/data?key=a", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: mockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	assertStatusCode(t, resp, http.StatusBadRequest)
	assertBody(t, resp, "no value set\n")
}
