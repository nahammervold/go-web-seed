package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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

	fmt.Printf("%s\n", body.Message)
}
