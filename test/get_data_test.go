package test

import (
	"net/http"
	"net/http/httptest"
	"simple-http-server/internal/endpoint"
	"testing"
)

func TestGetOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/data?key=a", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: mockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	assertStatusCode(t, resp, http.StatusOK)
	assertBody(t, resp, "AAA")
}

func TestGetBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/data?key=z", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: mockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	assertStatusCode(t, resp, http.StatusBadRequest)
	assertBody(t, resp, "no value for key z\n")
}
