package functions

import "regexp"

//HasNumbers rule
func HasNumbers(input string) bool {

	const pattern string = ".*[[:digit:]]"

	result, _ := regexp.MatchString(pattern, input)

	return result
}
