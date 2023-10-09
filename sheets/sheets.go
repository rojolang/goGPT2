// sheets.go
//	Author rjjm94 on 10/09/2023
// functions: ReadSettings
// variables: spreadsheetID
// description: This file contains the functionality to read settings from the Google Sheets.

package sheets

import (
	"github.com/rjjm94/gogpt2/config"
	"google.golang.org/api/sheets/v4"
	"strconv"
)

// ReadSettings reads the settings from the settings sheet and returns the settings.
func ReadSettings(srv *sheets.Service) (config.Settings, error) {
	// Read the settings from the settings sheet
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, "Settings!A1:B1000").Do()
	if err != nil {
		return config.Settings{}, err
	}

	// Parse the settings
	settings := config.Settings{}
	for _, row := range resp.Values {
		if len(row) < 2 {
			continue
		}

		key := row[0].(string)
		value := row[1].(string)

		switch key {
		case "PROMPT_SYSTEM_MESSAGE":
			settings.PromptSystemMessage = value
		case "PROMPT_TEMPERATURE":
			temp, _ := strconv.ParseFloat(value, 32)
			settings.PromptTemperature = float32(temp)
		case "PROMPT_MAX_TOKENS":
			settings.PromptMaxTokens, _ = strconv.Atoi(value)
		case "SHEET_NAME":
			settings.SheetName = value
		case "MESSAGE_VARIABLE1":
			settings.MessageVariable1 = value
		case "GENERATE_MESSAGE_VARIABLES_COLUMN":
			settings.GenerateMessageVariablesColumn = value
		case "PROMPT_COLUMN":
			settings.PromptColumn = value
		case "RESPONSE_COLUMN":
			settings.PromptResponseColumn = value
		case "VARIABLE1_COLUMN":
			settings.Variable1Column = value
		case "VARIABLE1_PROMPT_GENERATOR":
			settings.Variable1PromptSystemMessage = value
		case "VARIABLE1_PLACED_COLUMN":
			settings.Variable1PlacedColumn = value
		case "VARIABLE1_MAX_TOKENS":
			settings.Variable1MaxTokens, _ = strconv.Atoi(value)
		case "VARIABLE1_TEMPERATURE":
			temp, _ := strconv.ParseFloat(value, 32)
			settings.Variable1Temperature = float32(temp)
		}
	}

	return settings, nil
}
