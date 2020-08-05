# GoProof ( Alpha )

A simple and smooth validation package for GoLang ( and Frameworks )


Usage
---
```go
const myEmail string = "test@testgoproof.test"

validationErrors := Validate(myEmail, [][]string{
	{"email", "t", "it's not a valid email :( "},
})

if len(validationErrors) > 0 {
	log.Print(validationErrors)
}

```
**The "Validation" main function returns the errors.**




Functions and Keywords
---
```go

#-- For Strings


func MaxLength(stringInput string, stringValue string ), theMessage string) bool         ===> max
			
func MinLength(stringInput string, stringValue string ), theMessage string) bool         ===> min
			
func IsEmail(stringInput string), theMessage string) bool                                ===> email
			
func NotEmpty(stringInput string), theMessage string) bool                               ===> notempty
			
func IsURL(stringInput string), theMessage string) bool                                  ===> url
			
func IsIP(stringInput string), theMessage string) bool                                   ===> ip
			
func IsEqualString(stringInput string, stringValue string ), theMessage string) bool     ===> equal
			
func IsRgbHex(stringInput string), theMessage string) bool                               ===> rgbhex
			
func IsRgbDecimal(stringInput string), theMessage string) bool                           ===> rgb
			
func Uppercase(stringInput string), theMessage string) bool                              ===> upper
			
func Lowercase(stringInput string), theMessage string) bool                              ===> lower
			
func IsAlphaNumeric(stringInput string), theMessage string) bool                         ===> alphanumeric

func HasUppercase(stringInput string), theMessage string) bool                           ===> hasupper
		
func HasLowercase(stringInput string), theMessage string) bool                           ===> haslower
	
func HasNumbers(stringInput string), theMessage string) bool                             ===> hasnumber
	
func HasLetters(stringInput string), theMessage string) bool                             ===> hasletter
	
func ContainSomthing(stringInput string, stringValue string ), theMessage string) bool   ===> containsomthing




#-- For Numeric Values


func IsEqualNumeric(floatInput float64, floatValue float64), theMessage string) bool     ===> equal

func GreaterThan(floatInput float64, floatValue float64), theMessage string) bool        ===> gt

func GreaterThanOrEqual(floatInput float64, floatValue float64), theMessage string) bool ===> gteq

func LessThan(floatInput float64, floatValue float64), theMessage string) bool           ===> lt

func LessThanOrEqual(floatInput float64, floatValue float64), theMessage string) bool    ===> lteq



```