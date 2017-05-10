package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nahammervold/go-web-seed/models"
)

func TestSampleRoute(t *testing.T) {
	// Use the httptest struct to create a fake server
	ts := httptest.NewServer(http.HandlerFunc(SampleRoute))
	defer ts.Close()

	// Use the get request to get a response object or error
	res, err := http.Get(ts.URL)

	// Check for errors in the http Get function
	if err != nil {
		t.Fatal(err)
	}

	// Make sure Header has content-type set to json
	if res.Header.Get("Content-Type") != "application/json" {
		t.Fatal("Content-Type was not set to JSON")
	}

	var body models.SampleModel
	err = json.NewDecoder(res.Body).Decode(&body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	// Check for errors in res.Body
	if body.Message != "Hello world" {
		t.Fatal("Message doesn't contain correct content")
	}

	msg, err := http.Post(ts.URL, "application/json", strings.NewReader(`{"message": "hello"}`))

	if err != nil {
		t.Fatal(err)
	}

	if msg.StatusCode != http.StatusMethodNotAllowed {
		t.Fatal("Message should be rejected")
	}

}
