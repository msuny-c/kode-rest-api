package yaspeller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const URL = "https://speller.yandex.net/services/spellservice.json/checkText"

type Speller struct {
	config Config
	client http.Client
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

func NewWithConfig(config Config) *Speller {
	return &Speller{config: config, client: http.Client{}}
}

func New() *Speller {
	return NewWithConfig(Config{Lang:"ru,en", Options:"0", Format:"plain"})
}

func (speller *Speller) Text(text string) ([]Response, error) {
	formData := url.Values{}
	setupFormData(formData, speller, text)
	response, err := http.PostForm(URL, formData)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var data []Response
	json.Unmarshal(body, &data)
	return data, nil
}

func setupFormData(formData url.Values, speller *Speller, text string) {
	formData.Add("text", text)
	formData.Add("lang", speller.config.Lang)
	formData.Add("options", speller.config.Options)
	formData.Add("format", speller.config.Format)
}



