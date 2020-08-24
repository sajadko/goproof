package goproof

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/sajadko/goproof/functions"
)

//Validate | Main function for validating input
func Validate(input interface{}, rules [][]string, errorsArray *[]error) {
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
		floatValue, _ := strconv.ParseFloat(theValue, 64)

		if isInputStr {

			//Strings Rules
			switch theRule {

			//Required
			case "required":
				ResultSaver(functions.RequiredString(stringInput), theMessage)

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

			//Has Uppercase
			case "hasupper":
				ResultSaver(functions.HasUppercase(stringInput), theMessage)

			//Has Lowercase
			case "haslower":
				ResultSaver(functions.HasLowercase(stringInput), theMessage)

			//Has Number
			case "hasnumber":
				ResultSaver(functions.HasNumbers(stringInput), theMessage)

			//Has Letter
			case "hasletter":
				ResultSaver(functions.HasLetters(stringInput), theMessage)

			//Contain Somthing
			case "containsomthing":
				ResultSaver(functions.ContainSomthing(stringInput, stringValue), theMessage)

			//Is Time
			case "time":
				ResultSaver(functions.IsTime(stringInput, stringValue), theMessage)

			//Is Phone Number
			case "phone":
				ResultSaver(functions.IsPhoneNumber(stringInput), theMessage)

			//Is IMEI
			case "imei":
				ResultSaver(functions.IsIMEI(stringInput), theMessage)

			// //Is ISBN10
			// case "isbn10":
			// 	ResultSaver(functions.IsISBN10(stringInput), theMessage)

			// //Is ISBN13
			// case "isbn13":
			// 	ResultSaver(functions.IsISBN13(stringInput), theMessage)

			//Is IMEI
			case "json":
				ResultSaver(functions.IsJSON(stringInput), theMessage)

			//Is IMEI
			case "haswhitespace":
				ResultSaver(functions.HasWhiteSpace(stringInput), theMessage)

			//Is CreditCard
			case "creditcard":
				ResultSaver(functions.IsCreditCard(stringInput), theMessage)

			//Is DataURI
			case "datauri":
				ResultSaver(functions.IsDataURI(stringInput), theMessage)

				//!Test ISBN10
				//!Test ISBN13

				//Todo: IsASCII
				//Todo: Latitude
				//Todo: Longitude
				//Todo: IsDNSName
				//Todo: IsHexadecimal
				//Todo: IsHash (IsMD3 , MD4 , MD5 , IsSHA512 , IsSHA384 , IsSHA256)
			}
		} else if isInputFloat {

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

			//Multiple
			case "multiple":
				ResultSaver(functions.Multiple(floatInput, floatValue), theMessage)

			}

		}

	}

	for _, err := range theErrors {
		*errorsArray = append(*errorsArray, err)
	}

}

//Todo: Validate Struct
//Todo: Validate Map

//ValidateRequest validates (iris, gin, echo ...) framework requests and in a general form it can Validate any http.Request. this function take request and rules as input
func ValidateRequest(request *http.Request, rules map[string][][]string) map[string][]error {

	var errors map[string][]error = make(map[string][]error)
	var formValues map[string]string = make(map[string]string)

	r := request
	r.ParseMultipartForm(32)
	for key, val := range r.MultipartForm.Value {
		formValues[key] = val[0]
	}

	for key, val := range rules {
		var theErrors []error
		Validate(formValues[key], val, &theErrors)
		errors[key] = theErrors
	}

	fmt.Println(formValues)
	fmt.Println(errors)

	return errors

}

//=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/==> TESTING ROUTE (Just for test rules) <=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=/=

// func testRule(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")

// 	// const myEmail string = "test@testgoproof.test"

// 	// validationErrors := Validate(myEmail, [][]string{
// 	// 	{"email", "t", "it's not a valid email :( "},
// 	// })

// 	// if len(validationErrors) > 0 {
// 	// 	log.Print(validationErrors)
// 	// }

// 	var usernameErrors []error
// 	Validate("Sajadko1", [][]string{
// 		{"required", "t", "please enter a username"},
// 		{"max", "15", "max length for username is 15 character"},
// 		{"min", "7", "min length for username is 7 character"},
// 		{"hasupper", "t", "username must contain uppercase"},
// 		{"haslower", "t", "username must contain lowercase"},
// 		{"hasnumber", "t", "username must contain hasnumber"},
// 	}, &usernameErrors)

// 	var emailErrors []error
// 	Validate("sajadk48@gmail.com", [][]string{
// 		{"required", "t", "please enter an email"},
// 		{"email", "t", "please enter a valid email"},
// 	}, &emailErrors)

// 	var phoneErrors []error
// 	Validate("+79261234567", [][]string{
// 		{"phone", "t", "please enter an valid phone number"},
// 	}, &phoneErrors)

// 	var imeiErrors []error
// 	Validate("990000862471854", [][]string{
// 		{"imei", "t", "please enter an valid imei"},
// 	}, &imeiErrors)

// 	var errors = map[string][]error{
// 		"username": usernameErrors,
// 		"email":    emailErrors,
// 		"phone":    phoneErrors,
// 		"imei":     imeiErrors,
// 	}

// 	log.Println(errors)

// }

// func handleRequests() {
// 	http.HandleFunc("/", testRule)
// 	log.Fatal(http.ListenAndServe(":1998", nil))
// }

// func main() {
// 	handleRequests()
// }
