package models 

type Request struct {
	Note string `json:"note"`
}

type Response struct {
	Code int `json:"code,omitempty"`
	User string `json:"user,omitempty"`
	Note string `json:"created_note,omitempty"`
	Notes []string `json:"notes,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}

type Error struct {
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}