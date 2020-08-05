package functions

import "strconv"

//MaxLength | Max rule validation check
func MaxLength(input string, ruleValue string) bool {

	//Convert Rule Value to integer
	i, _ := strconv.ParseInt(ruleValue, 10, 64)

	return len(input) <= int(i)
}

//MinLength | Min rule validation check
func MinLength(input string, ruleValue string) bool {

	//Convert Rule Value to integer
	i, _ := strconv.ParseInt(ruleValue, 10, 64)

	return len(input) >= int(i)
}
