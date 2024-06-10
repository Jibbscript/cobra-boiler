package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCmd(t *testing.T) {
	// Set up the environment variable for the GitLab API token
	os.Setenv("GITLAB_TOKEN", "your_access_token")

	// Create a buffer to capture the output of the command
	var buf bytes.Buffer
	rootCmd.SetOutput(&buf)

	// Run the list command with the group ID and search term
	args := []string{"audit", "list", "-g", "6", "-s", "installation"}
	err := rootCmd.Execute(args)
	assert.NoError(t, err)

	// Check that the output contains the expected search results
	expectedOutput := `Filename: README.md
Path: README.md
Data: ```
## Installation

Quick start using the [pre-built
--------------------`
	assert.Contains(t, buf.String(), expectedOutput)
}