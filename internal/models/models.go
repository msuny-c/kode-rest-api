package models 

type Reponse struct {
	User string `json:"user,omitempty"`
	Note string `json:"created_note,omitempty"`
	Notes []string `json:"notes,omitempty"`
}

// type ResponseError struct {
// 	Code int `json:"code,omitempty"`
// 	Message string `json:"message,omitempty"`
// }

type Error struct {
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseError struct {
	Code int `json:"code,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}