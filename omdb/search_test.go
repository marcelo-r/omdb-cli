package omdb

import (
	"testing"
)

// TestRender checks if input and output are correct
func TestRender(t *testing.T) {
	response := ResponseSearch{
		Response: "true",
		Result: []SearchResult{
			{
				Poster: "poster_name.jpg",
				Title:  "The Test",
				Type:   "movie",
				Year:   "2021",
				Imdbid: "1234",
			},
		},
		TotalResults: "1",
	}
	got := response.Render()
	want := `
Found: 1
Showing the first 1

 
	Title: The Test (2021)
	Type: movie
	IMDb: https://www.imdb.com/title/1234

`
	if got != want {
		t.Fatalf("render does not match")
	}
}

func TestSearch(t *testing.T) {
	t.Fail()
}
