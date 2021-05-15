package omdb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// RESTResponse represents all action wich can be performed by a type of reponse
type RESTResponse interface {
	Load(*[]byte) error
	Render() string
}

const (
	// ApiKey used to access OMDb API
	ApiKey = "33063998"

	// baseURL is the basic template used in a request to OMDb API
	baseURL = "http://www.omdbapi.com/?apikey=%v"
)

// PrepareQuery takes a map of string:string and adds it as params and its
// values dest endpoint and setting the API Key to use
func PrepareQuery(params map[string]string) url.Values {
	query := url.Values{}
	for key, value := range params {
		query.Add(key, value)
	}
	return query
}

// Query queries the OMDB API for a given text
//
func Query(client *http.Client, endpoint RESTResponse, params map[string]string) error {
	// prepare request
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
		return fmt.Errorf("could not Do request: %w", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return fmt.Errorf("could not read response body: %w", err)
	}

	err = endpoint.Load(&data)
	// err = json.Unmarshal(data, &searchResult)
	if err != nil {
		return fmt.Errorf("could not load data into response: %w", err)
	}
	return nil
}
