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
	Errors []Error `json:"errors"`
}

type Error struct {
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

