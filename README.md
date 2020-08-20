# GoProof ( Alpha )

A simple and smooth validation package for GoLang ( and Frameworks )

![](https://img.shields.io/github/issues/sajadko/goproof?color=%23ff4800&style=for-the-badge)
![](https://img.shields.io/github/stars/sajadko/goproof?color=%23fe7f2d&style=for-the-badge)
![](https://img.shields.io/github/forks/sajadko/goproof?color=%23023e7d&style=for-the-badge)
![](https://img.shields.io/github/license/sajadko/goproof?color=%238ac926&style=for-the-badge)
<a href="https://pkg.go.dev/github.com/sajadko/goproof?tab=doc">
	<img src="https://img.shields.io/badge/pkg-Reference-00add8?logo=go&style=for-the-badge">
</a>

Usage Examples
---

#### Email Check
```go
const myEmail string = "test@testgoproof.test"

validationErrors := Validate(myEmail, [][]string{
	{"email", "t", "it's not a valid email :( "},
})

if len(validationErrors) > 0 {
	log.Print(validationErrors)
}

```

#### Number Check
```go
const number int = 25
validationErrors2 := Validate(number, [][]string{
	{"gt","12", "the number isn't greater than 12"},
	{"lt","30", "the number isn't less than 30"},
})

if len(validationErrors2) > 0 {
	log.Print(validationErrors2)
}

```

**The "Validation" main function returns an array of existing errors.**












Functions and Keywords
---
```go

#-- For Strings


func MaxLength(stringInput string, stringValue string ) bool         ===> max
			
func MinLength(stringInput string, stringValue string ) bool         ===> min
			
func IsEmail(stringInput string) bool                                ===> email
			
func NotEmpty(stringInput string) bool                               ===> notempty
			
func IsURL(stringInput string) bool                                  ===> url
			
func IsIP(stringInput string) bool                                   ===> ip
			
func IsEqualString(stringInput string, stringValue string ) bool     ===> equal
			
func IsRgbHex(stringInput string) bool                               ===> rgbhex
			
func IsRgbDecimal(stringInput string) bool                           ===> rgb
			
func Uppercase(stringInput string) bool                              ===> upper
			
func Lowercase(stringInput string) bool                              ===> lower
			
func IsAlphaNumeric(stringInput string) bool                         ===> alphanumeric

func HasUppercase(stringInput string) bool                           ===> hasupper
		
func HasLowercase(stringInput string) bool                           ===> haslower
	
func HasNumbers(stringInput string) bool                             ===> hasnumber
	
func HasLetters(stringInput string) bool                             ===> hasletter
	
func ContainSomthing(stringInput string, stringValue string ) bool   ===> containsomthing




#-- For Numeric Values


func IsEqualNumeric(floatInput float64, floatValue float64) bool     ===> equal

func GreaterThan(floatInput float64, floatValue float64) bool        ===> gt

func GreaterThanOrEqual(floatInput float64, floatValue float64) bool ===> gteq

func LessThan(floatInput float64, floatValue float64) bool           ===> lt

func LessThanOrEqual(floatInput float64, floatValue float64) bool    ===> lteq



```