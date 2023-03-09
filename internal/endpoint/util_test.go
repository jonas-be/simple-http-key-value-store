package endpoint

import (
	"io"
	"net/http"
	"simple-http-key-value-store/internal/database"
	"testing"
)

var MockDb = database.Database{
	"a": "AAA",
	"b": "bb",
	"c": "c",
}

func AssertDBEntry(t *testing.T, db Database, key string, expectedValue string) {
	if db.Get(key) != expectedValue {
		t.Errorf("expected \"%v\", got \"%s\"", expectedValue, db.Get(key))
	}
}

func AssertStatusCode(t *testing.T, resp *http.Response, expectedCode int) {
	if resp.StatusCode != expectedCode {
		t.Errorf("expected status %v, got %v", expectedCode, resp.Status)
	}
}

func AssertBody(t *testing.T, resp *http.Response, expectedBody string) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	s := string(body)

	if s != expectedBody {
		t.Errorf("expected \"%s\", got \"%s\"", expectedBody, s)
	}
}
