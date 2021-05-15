package omdb

import (
	"bytes"
	"encoding/json"
	"log"
	"text/template"
)

// ResponseSearch is the response without a type specified
type ResponseSearch struct {
	Response     string         `json:"Response"`
	Result       []SearchResult `json:"Search"`
	TotalResults string         `json:"totalResults"`
}

// SearchResult holds the response when search without a type
type SearchResult struct {
	Poster string `json:"Poster"`
	Title  string `json:"Title"`
	Type   string `json:"Type"`
	Year   string `json:"Year"`
	Imdbid string `json:"imdbID"`
}

var searchResultTemplate = `
Found: {{.TotalResults}}
Showing the first {{len .Result}}

{{range .Result}} 
	Title: {{.Title}} ({{.Year}})
	Type: {{.Type}}
	Id: {{.Imdbid}}
	IMDb: https://www.imdb.com/title/{{.Imdbid}}
{{end}}
`

// Render uses a template to render a reponse into a string
// used to render a response into a string for stdout
func (response *ResponseSearch) Render() string {
	searchTemplate, err := template.New("Search").Parse(searchResultTemplate)
	if err != nil {
		log.Fatal(err.Error())
	}
	var buffer bytes.Buffer
	if err = searchTemplate.Execute(&buffer, response); err != nil {
		log.Fatal(err.Error())
	}
	return buffer.String()
}

// Load serializes response data into a ResponseSearch
func (response *ResponseSearch) Load(data *[]byte) error {
	err := json.Unmarshal(*data, &response)
	return err
}
