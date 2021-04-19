package cmd

import (
	"fmt"
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
		searchFor := strings.Join(args, " ")
		fmt.Println("Looking up: ", searchFor)
		result := omdb.Search(&client, searchFor)
		fmt.Println(result.Render())
	},
}

func init() {
	rootCmd.Flags().BoolP("movie", "m", false, "Only movies")
	rootCmd.Flags().BoolP("series", "s", false, "Only series")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
