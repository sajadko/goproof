package functions

import (
	"encoding/json"
	"regexp"
)

//IsIMEI rule
func IsIMEI(input string) bool {

	const pattern string = "^[0-9a-f]{14}$|^\\d{15}$|^\\d{18}$"

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)
}

//IsPhoneNumber rule
func IsPhoneNumber(input string) bool {

	const pattern string = "^\\s*(?:\\+?(\\d{1,3}))?([-. (]*(\\d{3})[-. )]*)?((\\d{3})[-. ]*(\\d{2,4})(?:[-.x ]*(\\d+))?)\\s*$"

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)
}

// //IsISBN10 rule
// func IsISBN10(input string) bool {
// 	const pattern string = "^978(?:[0-9]{10})$"

// 	regex := regexp.MustCompile(pattern)

// 	return regex.MatchString(input)

// }

// //IsISBN13 rule
// func IsISBN13(input string) bool {

// 	const pattern string = "^978-(?:[0-9]{10})$"

// 	regex := regexp.MustCompile(pattern)

// 	return regex.MatchString(input)

// }

//IsJSON rule
func IsJSON(input string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(input), &js) == nil
}


//HasWhiteSpace rule
func HasWhiteSpace(input string) bool {

	const pattern string = ".*[[:space:]]"

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)

}

//IsCreditCard rule
func IsCreditCard(input string) bool {

	const pattern string = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11}|6[27][0-9]{14})$"

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)

}

//IsDataURI rule
func IsDataURI(input string) bool {

	const pattern string = "^data:(.+);base64,(.+)$"

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)

}