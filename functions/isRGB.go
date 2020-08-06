package functions

import "regexp"

//IsRgbDecimal rule
func IsRgbDecimal(input string) bool {
	const pattern string = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	result, _ := regexp.MatchString(pattern, input)
	return result
}

//IsRgbHex rule
func IsRgbHex(input string) bool {

	const pattern string = "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"

	result, _ := regexp.MatchString(pattern, input)

	return result
	
}
