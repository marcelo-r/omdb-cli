package omdb

import (
	"bytes"
	"encoding/json"
	"log"
	"text/template"
)

// ResponseMovie is the data returned when quering by type movie
type ResponseMovie struct {
	Actors     string   `json:"Actors"`
	Awards     string   `json:"Awards"`
	Boxoffice  string   `json:"BoxOffice"`
	Country    string   `json:"Country"`
	Dvd        string   `json:"DVD"`
	Director   string   `json:"Director"`
	Genre      string   `json:"Genre"`
	Language   string   `json:"Language"`
	Metascore  string   `json:"Metascore"`
	Plot       string   `json:"Plot"`
	Poster     string   `json:"Poster"`
	Production string   `json:"Production"`
	Rated      string   `json:"Rated"`
	Ratings    []rating `json:"Ratings"`
	Released   string   `json:"Released"`
	Response   string   `json:"Response"`
	Runtime    string   `json:"Runtime"`
	Title      string   `json:"Title"`
	Type       string   `json:"Type"`
	Website    string   `json:"Website"`
	Writer     string   `json:"Writer"`
	Year       string   `json:"Year"`
	Imdbid     string   `json:"imdbID"`
	Imdbrating string   `json:"imdbRating"`
	Imdbvotes  string   `json:"imdbVotes"`
}

type rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

var responseMovieTemplate = `
Title: {{.Title}} ({{.Country}}, {{.Year}}) - {{.Rated}} - |{{.Imdbrating}}|
Genre: {{.Genre}}

Director: {{.Director}}
Starring: {{.Actors}}

Plot: {{.Plot}}

`

// Load serializes response data into a ResponseSearch
func (response *ResponseMovie) Load(data *[]byte) error {
	err := json.Unmarshal(*data, &response)
	return err
}

func (response *ResponseMovie) Render() string {
	searchTemplate, err := template.New("Search").Parse(responseMovieTemplate)
	if err != nil {
		log.Fatal(err.Error())
	}
	var buffer bytes.Buffer
	if err = searchTemplate.Execute(&buffer, response); err != nil {
		log.Fatal(err.Error())
	}
	return buffer.String()
}
