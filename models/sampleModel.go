package models

// SampleModel structure used for responding in json
type SampleModel struct {
	// specify the json fieldname since Message needs to be capitalized in order to have visibility
	Message string `json:"message"`
}
