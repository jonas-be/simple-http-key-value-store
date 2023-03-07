package test

import (
	"io"
	"net/http"
	"simple-http-server/internal/database"
	"testing"
)

var mockDb = database.Database{
	"a": "AAA",
	"b": "bb",
	"c": "c",
}

func assertDBEntry(t *testing.T, db database.Database, key string, expectedValue string) {
	if db.Get(key) != expectedValue {
		t.Errorf("expected \"%v\", got \"%s\"", expectedValue, db.Get(key))
	}
}

func assertStatusCode(t *testing.T, resp *http.Response, expectedCode int) {
	if resp.StatusCode != expectedCode {
		t.Errorf("expected status %v, got %v", expectedCode, resp.Status)
	}
}

func assertBody(t *testing.T, resp *http.Response, expectedBody string) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	s := string(body)

	if s != expectedBody {
		t.Errorf("expected \"%s\", got \"%s\"", expectedBody, s)
	}
}
