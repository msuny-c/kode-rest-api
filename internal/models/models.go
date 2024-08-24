package models 

type Reponse struct {
	User string `json:"user,omitempty"`
	Note string `json:"created_note,omitempty"`
	Notes []string `json:"notes,omitempty"`
}

type ResponseError struct {
	Code int `json:"code"`
	Message string `json:"message"`
}