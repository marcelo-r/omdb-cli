package omdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// RESTResponse represents all action wich can be performed by a type of reponse
type RESTResponse interface {
	Load() error
	Render() string
	RenderHTML() string
}

const (
	// apiKey used to access OMDb API
	apiKey = "33063998"

	// baseURL is the basic template used in a request to OMDb API
	baseURL = "http://www.omdbapi.com/?apikey=%v"
)

// Search queries the OMDB API for a given text
func Search(client *http.Client, text string) RESTResponse {
	params := map[string]string{
		"apikey": apiKey,
		"s":      text,
	}
	request, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		fmt.Println("cant make request")
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	// mount query params
	query := PrepareQuery(params)
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}

	data, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	var searchResult ResponseSearch
	err = json.Unmarshal(data, &searchResult)
	if err != nil {
		log.Fatal(err.Error())
	}
	return searchResult
}

// PrepareQuery takes a map of string:string and adds it as params and its
// values dest endpoint and setting the API Key to use
func PrepareQuery(params map[string]string) url.Values {
	query := url.Values{}
	for key, value := range params {
		query.Add(key, value)
	}
	return query
}

type ResponseMovieTitle struct {
	Actors     string `json:"Actors"`
	Awards     string `json:"Awards"`
	Boxoffice  string `json:"BoxOffice"`
	Country    string `json:"Country"`
	Dvd        string `json:"DVD"`
	Director   string `json:"Director"`
	Genre      string `json:"Genre"`
	Language   string `json:"Language"`
	Metascore  string `json:"Metascore"`
	Plot       string `json:"Plot"`
	Poster     string `json:"Poster"`
	Production string `json:"Production"`
	Rated      string `json:"Rated"`
	Ratings    []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Released   string `json:"Released"`
	Response   string `json:"Response"`
	Runtime    string `json:"Runtime"`
	Title      string `json:"Title"`
	Type       string `json:"Type"`
	Website    string `json:"Website"`
	Writer     string `json:"Writer"`
	Year       string `json:"Year"`
	Imdbid     string `json:"imdbID"`
	Imdbrating string `json:"imdbRating"`
	Imdbvotes  string `json:"imdbVotes"`
}
