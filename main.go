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
