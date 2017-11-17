package models

// SampleModel structure used for responding in json
type SampleModel struct {

	// Unique Identifier
	ID string `json:"id"`

	// specify the json fieldname since Message needs to be capitalized in order to have visibility
	Message string `json:"message"`
}
