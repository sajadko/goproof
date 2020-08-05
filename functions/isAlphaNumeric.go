package functions

import "regexp"

//IsAlphaNumeric rule
func IsAlphaNumeric(input string) bool {

	const pattern string = "^[a-zA-Z0-9]+$"

	result, _ := regexp.MatchString(pattern, input)
	
	return result
	
}
