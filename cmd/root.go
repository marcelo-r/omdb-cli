package cmd

import (
	"fmt"
	"log"
	"net/http"
	"omdbcli/omdb"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// client should be the only http.Client used
var client http.Client = http.Client{}

var rootCmd = &cobra.Command{
	Use:   "omdb",
	Short: "query info about movies and tv shows from OMDb.com",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		searchText := strings.Join(args, " ")

		var searchParams map[string]string
		var endpointType omdb.RESTResponse

		if value, err := cmd.Flags().GetBool("series"); value && err == nil {
			searchParams = map[string]string{
				"apikey": omdb.ApiKey,
				"type":   "series",
				"t":      searchText,
			}
		} else if value, err := cmd.Flags().GetBool("movie"); value && err == nil {
			searchParams = map[string]string{
				"apikey": omdb.ApiKey,
				"type":   "movie",
				"t":      searchText,
			}
			endpointType = &omdb.ResponseMovie{}
		} else {
			searchParams = map[string]string{
				"apikey": omdb.ApiKey,
				"s":      searchText,
			}
			endpointType = &omdb.ResponseSearch{}
		}
		fmt.Println("Looking up: ", searchText)
		err := omdb.Query(&client, endpointType, searchParams)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(endpointType.Render())
	},
}

func init() {
	rootCmd.Flags().BoolP("movie", "m", false, "Only movies")
	rootCmd.Flags().BoolP("series", "s", false, "Only series")
}

// Execute the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
