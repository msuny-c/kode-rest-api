package yaspeller

import (
	"encoding/json"
	"io"
	"net/http"
)

var url = "https://speller.yandex.net/services/spellservice.json/checkText"

type Speller struct {
	config Config
}

type Config struct {
	Lang string
	Options string
	Format string
}

type Response struct {
	Code int `json:"code"`
	Pos int `json:"pos"`
	Row int `json:"row"`
	Col int `json:"col"`
	Len int `json:"len"`
	Word string `json:"word"`
}

func New() *Speller {
	return &Speller{Config{Lang:"ru,en", Options:"0", Format:"plain"}}
}

func NewWithConfig(config Config) *Speller {
	return &Speller{config: config}
}

func (speller *Speller) Text(text string) ([]Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	setupQuery(request, speller, text)
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var data []Response
	json.Unmarshal(body, &data)
	return data, nil
}

func setupQuery(request *http.Request, speller *Speller, text string) {
	query := request.URL.Query()
	query.Add("lang", speller.config.Lang)
	query.Add("options", speller.config.Options)
	query.Add("format", speller.config.Format)
	query.Add("text", text)
	request.URL.RawQuery = query.Encode()
}



