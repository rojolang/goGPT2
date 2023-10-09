// events.go
//  author: rjjm94 on 10/09/2023
// functions: triggerEvent, getMaxCols, getCellValues, printChangeMessage
// variables:
// description: This file contains the functionality to handle events.

package events

import (
	"fmt"
	"strings"
)

// triggerEvent triggers an event based on the updated cell's column
func triggerEvent(row, col int, prevValue, currValue interface{}) {
	switch indexToColumnLetter(col) {
	case "B":
		fmt.Printf("Event 1 triggered by change in cell at row %d column %d.\n", row, col)
		fmt.Printf("Previous value: %v, Current value: %v\n", prevValue, currValue)
	case "C":
		fmt.Printf("Event 2 triggered by change in cell at row %d column %d.\n", row, col)
		fmt.Printf("Previous value: %v, Current value: %v\n", prevValue, currValue)
	default:
		fmt.Printf("No event triggered for change in cell at row %d column %d.\n", row, col)
	}
}

// getMaxCols returns the maximum number of columns between the previous and current state of the row
func getMaxCols(row int, prevState, currState [][]interface{}) int {
	maxCols := 0
	if row < len(prevState) && row < len(currState) {
		maxCols = maxInt(len(prevState[row]), len(currState[row]))
	} else if row < len(prevState) {
		maxCols = len(prevState[row])
	} else if row < len(currState) {
		maxCols = len(currState[row])
	}
	return maxCols
}

// getCellValues returns the previous and current values of a cell
func getCellValues(row, col int, prevState, currState [][]interface{}) (interface{}, interface{}) {
	var prevValue, currValue interface{}
	if row < len(prevState) && col < len(prevState[row]) {
		prevValue = prevState[row][col]
	}
	if row < len(currState) && col < len(currState[row]) {
		currValue = currState[row][col]
	}
	return prevValue, currValue
}

// printChangeMessage prints a message about the change in a cell
func printChangeMessage(row, col int, prevValue, currValue interface{}) {
	if currValue != nil && strings.TrimSpace(currValue.(string)) != "" {
		if prevValue != nil {
			fmt.Printf("Cell at row %d column %d has changed from %v to %v\n", row+1, col+1, prevValue, currValue)
		} else {
			fmt.Printf("Cell at row %d column %d has been added with value %v\n", row+1, col+1, currValue)
		}
	} else if prevValue != nil && currValue == nil {
		fmt.Printf("Cell at row %d column %d has been cleared. Previous value was: %v\n", row+1, col+1, prevValue)
	}
}
