package functions

import "regexp"

//HasUppercase rule
func HasUppercase(input string) bool {

	const pattern string = ".*[[:upper:]]"

	result, _ := regexp.MatchString(pattern, input)

	return result

}

//HasLowercase rule
func HasLowercase(input string) bool {

	const pattern string = ".*[[:lower:]]"

	result, _ := regexp.MatchString(pattern, input)

	return result

}
