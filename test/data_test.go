package test

import (
	"net/http"
	"net/http/httptest"
	"simple-http-server/internal/endpoint"
	"simple-http-server/test/util"
	"testing"
)

func TestDataNoKey(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/data", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: util.MockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	util.AssertStatusCode(t, resp, http.StatusBadRequest)
	util.AssertBody(t, resp, "no key set")
}

func TestDataUnsupportedMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodOptions, "http://localhost:8080/data?key=a", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: util.MockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	util.AssertStatusCode(t, resp, http.StatusMethodNotAllowed)
	util.AssertBody(t, resp, "unsupported method")
}