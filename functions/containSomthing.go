package functions

import "strings"

//ContainSomthing rule (using strings standard package)
func ContainSomthing(input string, ruleValue string) bool {

	return strings.Contains(input, ruleValue)
	
}
