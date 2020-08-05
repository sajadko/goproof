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


func MaxLength(stringInput, stringValue), theMessage)           ===> max
			
func MinLength(stringInput, stringValue), theMessage)           ===> min
			
func IsEmail(stringInput), theMessage)                          ===> email
			
func NotEmpty(stringInput), theMessage                          ===> notempty
			
func IsURL(stringInput), theMessage)                            ===> url
			
func IsIP(stringInput), theMessage                              ===> ip
			
func IsEqualString(stringInput, stringValue), theMessage)       ===> equal
			
func IsRgbHex(stringInput), theMessage)                         ===> rgbhex
			
func IsRgbDecimal(stringInput), theMessage)                     ===> rgb
			
func Uppercase(stringInput), theMessage)                        ===> upper
			
func Lowercase(stringInput), theMessage)                        ===> lower
			
func IsAlphaNumeric(stringInput), theMessage)                   ===> alphanumeric

func HasUppercase(stringInput), theMessage                      ===> hasupper
		
func HasLowercase(stringInput), theMessage                      ===> haslower
	
func HasNumbers(stringInput), theMessage)                       ===> hasnumber
	
func HasLetters(stringInput), theMessage)                       ===> hasletter
	
func ContainSomthing(stringInput, stringValue), theMessage      ===> containsomthing




#-- For Numeric Values


func IsEqualNumeric(floatInput, floatValue), theMessage      	===> equal

func GreaterThan(floatInput, floatValue), theMessage      		===> gt

func GreaterThanOrEqual(floatInput, floatValue), theMessage     ===> gteq

func LessThan(floatInput, floatValue), theMessage      			===> lt

func LessThanOrEqual(floatInput, floatValue), theMessage      	===> lteq



```