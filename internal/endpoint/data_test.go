package endpoint

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestDataNoKey(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/data", nil)
	w := httptest.NewRecorder()

	db := NewMockDatabase(t)

	dataHandler := DataHandler{Db: db}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusBadRequest)
	AssertBody(t, resp, "no key set\n")
}

func TestDataUnsupportedMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodOptions, "http://localhost:8080/data?key=a", nil)
	w := httptest.NewRecorder()

	db := NewMockDatabase(t)

	dataHandler := DataHandler{Db: db}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusMethodNotAllowed)
	AssertBody(t, resp, "unsupported method\n")
}

func TestGetOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/data?key=a", nil)
	w := httptest.NewRecorder()

	db := NewMockDatabase(t)
	db.EXPECT().Contains("a").Return(true).Once()
	db.EXPECT().Get("a").Return("AAA").Once()

	dataHandler := DataHandler{Db: db}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusOK)
	AssertBody(t, resp, "AAA\n")
}

func TestGetBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/data?key=z", nil)
	w := httptest.NewRecorder()

	db := NewMockDatabase(t)
	db.EXPECT().Contains("z").Return(false).Once()

	dataHandler := DataHandler{Db: db}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusBadRequest)
	AssertBody(t, resp, "no value for key z\n")
}

func TestPutOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/data?key=b&value=ZZZ", nil)
	w := httptest.NewRecorder()

	db := NewMockDatabase(t)
	db.EXPECT().Contains("b").Return(true).Once()
	db.EXPECT().Set("b", "ZZZ").Return(nil).Once()

	dataHandler := DataHandler{Db: db}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusOK)
	AssertBody(t, resp, "\n")
}

func TestPutCreated(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/data?key=x&value=X", nil)
	w := httptest.NewRecorder()

	db := NewMockDatabase(t)
	db.EXPECT().Contains("x").Return(false).Once()
	db.EXPECT().Set("x", "X").Return(nil).Once()

	dataHandler := DataHandler{Db: db}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusCreated)
	AssertBody(t, resp, "\n")
}

func TestPutBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/data?key=a", nil)
	w := httptest.NewRecorder()

	db := NewMockDatabase(t)

	dataHandler := DataHandler{Db: db}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusBadRequest)
	AssertBody(t, resp, "no value set\n")
}

func TestDeleteOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/data?key=c", nil)
	w := httptest.NewRecorder()

	db := NewMockDatabase(t)
	db.EXPECT().Contains("c").Return(true).Once()
	db.EXPECT().Delete("c").Return().Once()

	dataHandler := DataHandler{Db: db}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusNoContent)
	AssertBody(t, resp, "\n")
}

func TestDeleteBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/data?key=z", nil)
	w := httptest.NewRecorder()

	db := NewMockDatabase(t)
	db.EXPECT().Contains("z").Return(false).Once()

	dataHandler := DataHandler{Db: db}
	dataHandler.HandleRequest(w, req)

	resp := w.Result()

	AssertStatusCode(t, resp, http.StatusBadRequest)
	AssertBody(t, resp, "key z not exists\n")
}

func TestRaceCondition(t *testing.T) {
	db := NewMockDatabase(t)
	db.EXPECT().Contains("x").Return(true).Once()
	db.EXPECT().Set("x", "AAA").Return(nil).Once()

	dataHandler := DataHandler{Db: db}

	var wg sync.WaitGroup
	wg.Add(2)

	go call(dataHandler, http.MethodGet, "http://localhost:8080/data?key=x", &wg)
	go call(dataHandler, http.MethodPut, "http://localhost:8080/data?key=x&value=AAA", &wg)

	wg.Wait()
}

func call(dataHandler DataHandler, method string, url string, wg *sync.WaitGroup) {
	defer wg.Done()

	req := httptest.NewRequest(method, url, nil)
	w := httptest.NewRecorder()

	dataHandler.HandleRequest(w, req)

	fmt.Printf("%v: %v", url, w.Result().StatusCode)
}
