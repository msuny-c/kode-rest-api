package models 

type Request struct {
	Note string `json:"note"`
}

type ResponseNote struct {
	Note string `json:"created_note"`
}

type ResponseNotes struct {
	Notes []string `json:"notes"`
}

type ResponseError struct {
	Errors Errors `json:"errors"`
}

type Errors []Error

type Error struct {
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (errors Errors) Add(code string, message string) {
	errors = append(errors, Error{Code: code, Message: message})
}

