package test

import (
	"net/http"
	"net/http/httptest"
	"simple-http-server/internal/endpoint"
	"simple-http-server/test/util"
	"testing"
)

func TestGetOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/data?key=a", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: util.MockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	util.AssertStatusCode(t, resp, http.StatusOK)
	util.AssertBody(t, resp, "AAA")
}

func TestGetBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/data?key=z", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: util.MockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	util.AssertStatusCode(t, resp, http.StatusBadRequest)
	util.AssertBody(t, resp, "no value for key z\n")
}
