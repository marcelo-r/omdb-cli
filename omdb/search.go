package omdb

import (
	"bytes"
	"log"
	"text/template"
)

// ResponseSearch is the response without a type specified
type ResponseSearch struct {
	Response     string         `json:"Response"`
	Result       []SearchResult `json:"Search"`
	TotalResults string         `json:"totalResults"`
}

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
	IMDb: https://www.imdb.com/title/{{.Imdbid}}
{{end}}
`

// Render uses a template to render a reponse into a string
func (response ResponseSearch) Render() string {
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

func (response ResponseSearch) Load() error {
	return nil
}

func (response ResponseSearch) RenderHTML() string {
	return ""
}
