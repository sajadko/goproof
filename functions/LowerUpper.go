package functions

import (
	"log"
	"regexp"
)

//Lowercase |
func Lowercase(input string) bool {
	result, err := regexp.MatchString("^[a-z]+$", input)

	if err != nil {
		log.Panic(err)
	}

	return result
}

//Uppercase |
func Uppercase(input string) bool {
	result, err := regexp.MatchString("^[A-Z]+$", input)

	if err != nil {
		log.Panic(err)
	}

	return result
}
