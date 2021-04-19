package main

import "omdbcli/cmd"

var searchTypes = map[string]string{
	"movie":  "type=movie",
	"series": "type=series",
	"any":    "",
}

func main() {
	cmd.Execute()
}

//func main() {
//	url := fmt.Sprintf(baseURL, apiKey)
//	url += "&type=movie&t=matrix"
//	client := &http.Client{}
//	resp, err := client.Get(url)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	fmt.Printf("%s", body)
//}
