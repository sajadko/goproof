package functions

//Multiple rule
func Multiple(input float64, ruleValue float64) bool {

	inputint := int64(input)
	valuetint := int64(ruleValue)

	return inputint%valuetint == 0

}
