package main

import (
	"GoProof/functions"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//Validate | Main function for validating input
func Validate(input interface{}, rules [][]string) []error {
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

	ExtractValues := func(rule []string) (string, string, string) {
		var theRule string
		var theValue string
		var theMessage string = ""

		theRule = rule[0]
		theValue = rule[1]

		//Extract custom message from 'message'
		messageRemovedLeftQ := strings.TrimLeft(rule[2], "'")
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

			//Is Equal String
			case "equal":
				ResultSaver(functions.IsEqualString(stringInput, stringValue), theMessage)

			//Is RGB Hex
			case "rgbhex":
				ResultSaver(functions.IsRgbHex(stringInput), theMessage)

			//Is RGB (Decimal)
			case "rgb":
				ResultSaver(functions.IsRgbDecimal(stringInput), theMessage)

			//Is Uppercase
			case "upper":
				ResultSaver(functions.Uppercase(stringInput), theMessage)

			//Is Lowercase
			case "lower":
				ResultSaver(functions.Lowercase(stringInput), theMessage)

			//isAlphaNumeric
			case "alphanumeric":
				ResultSaver(functions.IsAlphaNumeric(stringInput), theMessage)

			//hasUppercase
			case "hasupper":
				ResultSaver(functions.HasUppercase(stringInput), theMessage)

			//isAlphaNumeric
			case "haslower":
				ResultSaver(functions.HasLowercase(stringInput), theMessage)

			case "hasnumber":
				ResultSaver(functions.HasNumbers(stringInput), theMessage)

			case "hasletter":
				ResultSaver(functions.HasLetters(stringInput), theMessage)

			case "containsomthing":
				ResultSaver(functions.ContainSomthing(stringInput, stringValue), theMessage)

			}
		} else if isInputFloat && isValueNumeric {

			//Numeric Rules
			switch theRule {

			//Is Equal Numeric
			case "equal":
				ResultSaver(functions.IsEqualNumeric(floatInput, floatValue), theMessage)

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

//=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/==> TESTING ROUTE (Just for test rules) <=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=

func testRule(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")

	const myEmail string = "test@testgoproof.test"

	validationErrors := Validate(myEmail, [][]string{
		{"email", "t", "it's not a valid email :( "},
	})

	if len(validationErrors) > 0 {
		log.Print(validationErrors)
	}


	const number int = 25
	validationErrors2 := Validate(number, [][]string{
		{"gt","12", "the number isn't greater than 12"},
		{"lt","30", "the number isn't less than 30"},
	})
	
	if len(validationErrors2) > 0 {
		log.Print(validationErrors2)
	}


}

func handleRequests() {
	http.HandleFunc("/", testRule)
	log.Fatal(http.ListenAndServe(":1998", nil))
}

func main() {
	handleRequests()
}
