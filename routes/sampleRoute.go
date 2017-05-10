package routes

import (
	"encoding/json"
	"net/http"

	"github.com/nahammervold/go-web-seed/models"
)

// SampleRoute Sample route that simply responds back with Hello World
func SampleRoute(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		// set header to json header
		res.Header().Set("Content-Type", "application/json")

		// create a response message
		msg := models.SampleModel{
			Message: "Hello world",
		}

		// send response back in JSON encoding
		json.NewEncoder(res).Encode(msg)
		break
	default:
		http.Error(res, "Only Get Requests are allowed", http.StatusMethodNotAllowed)
		break
	}

}
