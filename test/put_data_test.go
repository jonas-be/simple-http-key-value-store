package test

import (
	"net/http"
	"net/http/httptest"
	"simple-http-server/internal/endpoint"
	"simple-http-server/test/util"
	"testing"
)

func TestPutOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/data?key=b&value=ZZZ", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: util.MockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	util.AssertStatusCode(t, resp, http.StatusOK)
	util.AssertBody(t, resp, "")

	util.AssertDBEntry(t, dataHandler.Db, "b", "ZZZ")
}

func TestPutCreated(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/data?key=x&value=X", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: util.MockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	util.AssertStatusCode(t, resp, http.StatusCreated)
	util.AssertBody(t, resp, "")

	util.AssertDBEntry(t, dataHandler.Db, "x", "X")
}

func TestPutBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/data?key=a", nil)
	w := httptest.NewRecorder()

	dataHandler := endpoint.DataHandler{Db: util.MockDb}
	dataHandler.HandelRequest(w, req)

	resp := w.Result()

	util.AssertStatusCode(t, resp, http.StatusBadRequest)
	util.AssertBody(t, resp, "no value set\n")
}
