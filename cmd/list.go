/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the group ID from the command line arguments
		groupId, err := cmd.Flags().GetString("group-id")
		if err != nil {
			fmt.Println("Error getting group ID:", err)
			os.Exit(1)
		}

		// Get the search term from the command line arguments
		searchTerm, err := cmd.Flags().GetString("search-term")
		if err != nil {
			fmt.Println("Error getting search term:", err)
			os.Exit(1)
		}

		// Get the GitLab API token from the environment variable
		gitlabToken := os.Getenv("GITLAB_TOKEN")
		if gitlabToken == "" {
			fmt.Println("Error: GITLAB_TOKEN environment variable not set.")
			os.Exit(1)
		}

		// Construct the GitLab API URL
		gitlabApiUrl := fmt.Sprintf("https://gitlab.example.com/api/v4/groups/%s/search?scope=blobs&search=%s", groupId, searchTerm)

		// Create a new HTTP request
		req, err := http.NewRequest(http.MethodGet, gitlabApiUrl, nil)
		if err != nil {
			fmt.Println("Error creating HTTP request:", err)
			os.Exit(1)
		}

		// Set the Authorization header with the GitLab API token
		req.Header.Set("PRIVATE-TOKEN", gitlabToken)

		// Send the HTTP request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending HTTP request:", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			os.Exit(1)
		}

		// Parse the JSON response
		var results []map[string]interface{}
		err = json.Unmarshal(body, &results)
		if err != nil {
			fmt.Println("Error parsing JSON response:", err)
			os.Exit(1)
		}

		// Print the search results
		for _, result := range results {
			fmt.Println("Filename:", result["filename"])
			fmt.Println("Path:", result["path"])
			fmt.Println("Data:", result["data"])
			fmt.Println("--------------------")
		}
	},
}

func init() {
	auditCmd.AddCommand(listCmd)

}
