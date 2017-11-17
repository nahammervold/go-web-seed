package routes

import (
	"encoding/json"
	"net/http"

	"github.com/nahammervold/go-web-seed/models"
	"github.com/satori/go.uuid"
)

type Route interface {
	Get (res http.ResponseWriter, req *http.Request)
	Post (res http.ResponseWriter, req *http.Request)
	Put (res http.ResponseWriter, req *http.Request)
	Delete (res http.ResponseWriter, req *http.Request)
}

type sample struct {
	msg []models.SampleModel
}

func (s *sample) Get (response http.ResponseWriter, req *http.Request){

	// set header to json header
	response.Header().Set("Content-Type", "application/json")

	json.NewEncoder(response).Encode(s.msg);
}

func (s *sample) Post (res http.ResponseWriter, req *http.Request){

	// set header to json header
	res.Header().Set("Content-Type", "application/json")

	// grab body from request and populate a SampleModel
	msg := models.SampleModel{
		ID: uuid.NewV1().String(),
	}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&msg)

	if err != nil {
		http.Error(res, "Could not create", http.StatusForbidden)
		return
	}

	// add to message slice
	s.msg = append(s.msg, msg)

	// respond to request with original message
	json.NewEncoder(res).Encode(msg)
}

func (s *sample) Put (res http.ResponseWriter, req *http.Request){
	found := false

	// set header to json header
	res.Header().Set("Content-Type", "application/json")

	// grab body from request and populate a SampleModel
	msg := models.SampleModel{

	}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&msg)

	if err != nil {
		http.Error(res, "Could not Delete", http.StatusForbidden)
		return
	}

	// update existing element with matching id
	for i := range s.msg {
		if s.msg[i].ID == msg.ID {
			s.msg[i].Message = msg.Message
			found = true
		}
	}

	if found {
		// respond to request with original message
		json.NewEncoder(res).Encode(msg)
	} else {
		http.Error(res, "Could not Update", http.StatusForbidden)
	}

}

func (s *sample) Delete(res http.ResponseWriter, req *http.Request){
	found := false

	// set header to json header
	res.Header().Set("Content-Type", "application/json")

	// grab body from request and populate a SampleModel
	msg := models.SampleModel{

	}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&msg)

	if err != nil {
		http.Error(res, "Could not Delete", http.StatusForbidden)
		return
	}

	// delete existing element with matching id
	for i := range s.msg {
		if s.msg[i].ID == msg.ID {
			found = true
			copy(s.msg[i:], s.msg[i+1:])
			s.msg[len(s.msg)-1] = models.SampleModel{} // or the zero value of T
			s.msg = s.msg[:len(s.msg)-1]
			break
		}
	}

	if found {
		// respond to request with original message
		json.NewEncoder(res).Encode(s.msg)
	} else {
		http.Error(res, "Could not Delete", http.StatusForbidden)
	}
}

var s sample

// SampleRoute Sample route that simply responds back with Hello World
func SampleRoute(res http.ResponseWriter, req *http.Request) {

	if s.msg == nil {
		s.msg = make([]models.SampleModel, 0, 10)
	}

	switch req.Method {
	case http.MethodGet:
		s.Get(res, req)
		break
	case http.MethodPost:
		s.Post(res,req)
		break
	case http.MethodPut:
		s.Put(res, req)
		break
	case http.MethodPatch:
		s.Put(res, req)
		break
	case http.MethodDelete:
		s.Delete(res, req)
		break
	default:
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		break
	}

}
