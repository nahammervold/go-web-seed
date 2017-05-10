package routes

import (
	"net/http"
)

// NewRouter Returns a ServeMux pointer with the routes attached to it
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/sample", SampleRoute)

	return mux
}
