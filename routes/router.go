package routes

import (
	"net/http"
)

// GetRouter Returns a ServeMux pointer with the routes attached to it
func GetRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/sample", SampleRoute)

	return mux
}
