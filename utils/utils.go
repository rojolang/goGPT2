// utils.go
// Created by rjjm94 on 10/09/2023
// functions: contains, isValidVariable1Value, columnLetterToIndex, indexToColumnLetter
// variables: maxInt
// description: This file contains utility functions.

package utils

import "fmt"

// MaxInt returns the maximum of two integers
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Contains checks if a slice contains a value
func Contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// IsValidVariable1Value checks if the value is a valid variable1 value
func IsValidVariable1Value(value interface{}) bool {
	validValues := []string{"POSITIVE", "NEGATIVE", "NUETRAL", "SHOCKING", "TRENDING", "HOPEFUL"}
	for _, validValue := range validValues {
		if value == validValue {
			return true
		}
	}
	return false
}

// ColumnLetterToIndex converts a column letter to a column index
func ColumnLetterToIndex(letter string) int {
	index := 0
	multiplier := 1
	for i := len(letter) - 1; i >= 0; i-- {
		char := letter[i]
		if char >= 'a' && char <= 'z' {
			char -= 'a' - 'A'
		}
		index += int(char-'A'+1) * multiplier
		multiplier *= 26
	}
	return index - 1
}

// IndexToColumnLetter converts a column index to a column letter
func IndexToColumnLetter(index int) string {
	letter := ""
	for index >= 0 {
		letter = fmt.Sprintf("%c", 'A'+index%26) + letter
		index = index/26 - 1
	}
	return letter
}
