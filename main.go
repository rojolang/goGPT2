// main.go
// author: rjjm94 on 10/09/2023
// functions: startMainLoop
// variables: prevState
// description: This file contains the main loop of the application.

package main

import (
	"log"

	"github.com/rjjm94/gogpt2/config"
	"github.com/rjjm94/gogpt2/sheets"
)

func main() {
	// SetupEnvironment loads environment variables, creates the Google Sheets and Redis services.
	ctx, redisClient, srv := config.SetupEnvironment()

	// Read the settings from the settings sheet.
	settings, err := sheets.ReadSettings(srv)
	if err != nil {
		log.Fatalf("Error reading settings: %v", err)
	}

	// Initialize the previous state of the Google Sheet
	var prevState [][]interface{}

	// Start the main loop
	for {
		startMainLoop(ctx, redisClient, srv, settings, &prevState)
	}
}
