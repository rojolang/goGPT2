// loop.go
// author: rjjm94 on 10/09/2023
// functions: StartMainLoop, handleCellValueChange
// variables: minPromptLen
// description: This file contains the functionality to start the main loop of the application.

package loop

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/rjjm94/gogpt2/sheets"
)

// StartMainLoop starts the main loop of the application.
func StartMainLoop(srv *sheets.Service, spreadsheetID string) {
	for {
		// Read the settings
		settings, err := sheets.ReadSettings(srv)
		if err != nil {
			// Handle error...
		}

		// Get the cell values
		resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, settings.SheetName+"!A1:Z1000").Do()
		if err != nil {
			// Handle error...
		}

		// Handle cell value changes
		for _, row := range resp.Values {
			// Your existing code...
		}

		// Sleep for a while before the next iteration
		time.Sleep(time.Second * 10)
	}
}

// handleCellValueChange handles a change in a cell value.
func handleCellValueChange(ctx context.Context, srv *sheets.Service, settings Settings, i int, j int, prompt interface{}, wg *sync.WaitGroup, ch chan struct{}, prevState *[][]interface{}, resp *sheets.ValueRange) {
	// Decrement the WaitGroup counter when the goroutine completes
	defer wg.Done()

	// Acquire a token from the channel
	ch <- struct{}{}
	// Release the token when the goroutine completes
	defer func() { <-ch }()

	// Convert the prompt to a string and check if it's at least 20 characters long
	if promptStr, ok := prompt.(string); ok && len(promptStr) >= minPromptLen {
		systemMessage := settings.PromptSystemMessage
		maxTokens := settings.PromptMaxTokens
		temperature := settings.PromptTemperature

		// If the cell is in the prompt column, generate a response using the GPT-4 API and update the cell in the response column
		if indexToColumnLetter(j) == settings.PromptColumn {
			// Check if the variable1 column has a valid value
			variable1Value, _ := getCellValues(i, columnLetterToIndex(settings.Variable1Column), *prevState, resp.Values)
			if !isValidVariable1Value(variable1Value) {
				// If not, write a message to the variable1 column and skip this iteration
				err := updateSheetWithBackoff(ctx, srv, spreadsheetID, fmt.Sprintf("%s!%s%d", settings.SheetName, settings.Variable1PlacedColumn, i+1), "PLEASE FILL IN")
				if err != nil {
					log.Printf("Error updating cell in Google Sheet for row %d: %v", i+1, err)
				}
				return
			}

			// Generate a response using the GPT-4 API
			response, err := getGPT4Response(promptStr, systemMessage, maxTokens, float32(temperature))
			if err != nil {
				log.Printf("Error getting response from GPT-4 API for row %d: %v", i+1, err)
				return
			}

			// Update the cell in the response column
			err = updateSheetWithBackoff(ctx, srv, spreadsheetID, fmt.Sprintf("%s!%s%d", settings.SheetName, settings.PromptResponseColumn, i+1), response)
			if err != nil {
				log.Printf("Error updating cell in Google Sheet for row %d: %v", i+1, err)
				return
			}

			log.Printf("Successfully updated cell in Google Sheet for row %d", i+1)
		} else if indexToColumnLetter(j) == settings.Variable1Column {
			// If the cell is in the variable column, generate a response using the variable prompt generator and update the cell in the generate message variables column
			variableMaxTokens := settings.Variable1MaxTokens
			variableTemperature := settings.Variable1Temperature

			// Generate a response using the variable prompt generator
			response, err := getGPT4Response(promptStr, settings.Variable1PromptSystemMessage, variableMaxTokens, float32(variableTemperature))
			if err != nil {
				log.Printf("Error getting response from GPT-4 API for row %d: %v", i+1, err)
				return
			}

			// Update the cell in the generate message variables column
			err = updateSheetWithBackoff(ctx, srv, spreadsheetID, fmt.Sprintf("%s!%s%d", settings.SheetName, settings.Variable1PlacedColumn, i+1), response)
			if err != nil {
				log.Printf("Error updating cell in Google Sheet for row %d: %v", i+1, err)
				return
			}

			log.Printf("Successfully updated cell in Google Sheet for row %d", i+1)
		}
	}
}
