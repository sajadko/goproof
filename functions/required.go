package functions

import "strings"

//RequiredString rule
func RequiredString(input string) bool {
	
	return strings.TrimSpace(input) != ""
	
}

//OptionalString rule
func OptionalString(input string) bool {
	
	

	return false
}

//RequiredNumeric rule
func RequiredNumeric(input float64) bool {
	return false
}

//OptionalNumeric rule
func OptionalNumeric(input float64) bool {
	return false
}
