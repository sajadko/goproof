package main

import (
	"GoProof/functions"
	"errors"
	"strconv"
	"strings"
)

//Validate | Main function for validating input
func Validate(input interface{}, rules []string) []error {
	var theErrors []error

	//ResultSaver | This function pushes the result to Errors array
	ResultSaver := func(res bool, message string) {
		var theError error
		if !res {
			theError = errors.New(message)
		} else {
			theError = errors.New("")
		}
		if len(theError.Error()) > 0 {
			theErrors = append(theErrors, theError)
		}
	}

	ExtractValues := func(rule string) (string, string, string) {
		var theRule string
		var theValue string
		var theMessage string = ""

		ruleAndMessage := strings.Split(rule, "~>")
		ruleAndValue := strings.Split(ruleAndMessage[0], "=")

		//Extract value from 'value'
		valueWithoutQ := strings.Split(ruleAndValue[1], "'")

		//Set Rule and value
		theRule = ruleAndValue[0]
		theValue = valueWithoutQ[1]

		//Extract custom message from 'message'
		messageRemovedLeftQ := strings.TrimLeft(ruleAndMessage[1], "'")
		messageRemovedRightQ := strings.TrimRight(messageRemovedLeftQ, "'")
		messageWithoutQ := messageRemovedRightQ

		//Set custom message if exists
		theMessage = messageWithoutQ

		return theRule, theValue, theMessage
	}

	//*Check input is string or number(float64)
	stringInput, isInputStr := input.(string)
	floatInput, isInputFloat := input.(float64)
	intInput, isInputInt := input.(int)
	if isInputInt {
		floatInput = float64(intInput)
		isInputFloat = true
	}

	// log.Print(isInputStr, isInputFloat, isInputInt)

	//Iterate over all rules
	for _, rule := range rules {

		theRule, theValue, theMessage := ExtractValues(rule)
		// log.Print(theRule, theValue, theMessage, "\n")

		//*Check value is string or number(float64)
		stringValue := theValue
		floatValue, errFloat := strconv.ParseFloat(theValue, 64)
		var isValueNumeric bool
		if errFloat != nil {
			isValueNumeric = false
		} else {
			isValueNumeric = true
		}

		if isInputStr && !isValueNumeric {

			//Strings Rules
			switch theRule {
			//Maximum Length
			case "max":
				ResultSaver(functions.MaxLength(stringInput, stringValue), theMessage)

			//Minimum Length
			case "min":
				ResultSaver(functions.MinLength(stringInput, stringValue), theMessage)

			//Is Email
			case "email":
				ResultSaver(functions.IsEmail(stringInput), theMessage)

			//Not Empty
			case "notempty":
				ResultSaver(functions.NotEmpty(stringInput), theMessage)

			//Is URL
			case "url":
				ResultSaver(functions.IsURL(stringInput), theMessage)

			case "ip":
				ResultSaver(functions.IsIP(stringInput), theMessage)

			case "cupper":

				//Todo: String Compare Equality
				//Todo: Contain Uppercase
				//Todo: Contain Lowercase
				//Todo: Contain Numbers
				//Todo: Contain Alphabet
				//Todo: isAlphaNumeric
				//Todo: Contain somthing
				//Todo: Uppercase & Lowercase
				//Todo: RGB HEX & RGB Decimal
			}
		} else if isInputFloat && isValueNumeric {

			//Numeric Rules
			switch theRule {

			//Greater Than
			case "gt":
				ResultSaver(functions.GreaterThan(floatInput, floatValue), theMessage)

			//Greater Than Or Equal
			case "gteq":
				ResultSaver(functions.GreaterThanOrEqual(floatInput, floatValue), theMessage)

			//Less Than
			case "lt":
				ResultSaver(functions.LessThan(floatInput, floatValue), theMessage)

			//Less Than Or Equal
			case "lteq":
				ResultSaver(functions.LessThanOrEqual(floatInput, floatValue), theMessage)

			}

		}

	}

	return theErrors
}
