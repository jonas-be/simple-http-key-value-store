package endpoint

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDataNoKey(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/data", nil)
	w := httptest.NewRecorder()

	dataHandler := DataHandler{Db: MockDb}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusBadRequest)
	AssertBody(t, resp, "no key set\n")
}

func TestDataUnsupportedMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodOptions, "http://localhost:8080/data?key=a", nil)
	w := httptest.NewRecorder()

	dataHandler := DataHandler{Db: MockDb}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusMethodNotAllowed)
	AssertBody(t, resp, "unsupported method\n")
}

func TestGetOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/data?key=a", nil)
	w := httptest.NewRecorder()

	dataHandler := DataHandler{Db: MockDb}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusOK)
	AssertBody(t, resp, "AAA")
}

func TestGetBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/data?key=z", nil)
	w := httptest.NewRecorder()

	dataHandler := DataHandler{Db: MockDb}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusBadRequest)
	AssertBody(t, resp, "no value for key z\n")
}

func TestPutOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/data?key=b&value=ZZZ", nil)
	w := httptest.NewRecorder()

	dataHandler := DataHandler{Db: MockDb}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusOK)
	AssertBody(t, resp, "")

	AssertDBEntry(t, dataHandler.Db, "b", "ZZZ")
}

func TestPutCreated(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/data?key=x&value=X", nil)
	w := httptest.NewRecorder()

	dataHandler := DataHandler{Db: MockDb}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusCreated)
	AssertBody(t, resp, "")

	AssertDBEntry(t, dataHandler.Db, "x", "X")
}

func TestPutBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/data?key=a", nil)
	w := httptest.NewRecorder()

	dataHandler := DataHandler{Db: MockDb}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusBadRequest)
	AssertBody(t, resp, "no value set\n")
}

func TestDeleteOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/data?key=c", nil)
	w := httptest.NewRecorder()

	dataHandler := DataHandler{Db: MockDb}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusOK)
	AssertBody(t, resp, "")

	if dataHandler.Db.Contains("c") {
		t.Error("expected \"c\" not exists, got \"c\" exists")
	}
}

func TestDeleteBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/data?key=z", nil)
	w := httptest.NewRecorder()

	dataHandler := DataHandler{Db: MockDb}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusBadRequest)
	AssertBody(t, resp, "key z not exists\n")
}
