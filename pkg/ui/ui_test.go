package ui

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestWebHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	// override STATIC_FILES_PATH as test runs from pkg
	os.Setenv("STATIC_FILES_PATH", "../../web/")
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("error: ", err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(WebHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	f, err := os.Open("../../web/index.html")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	expected, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
}
