package functions

import "regexp"

//HasLetters rule
func HasLetters(input string) bool {

	const pattern string = ".*[[:alpha:]]"

	result, _ := regexp.MatchString(pattern, input)

	return result
}