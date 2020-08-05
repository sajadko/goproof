package functions

import "strings"

//NotEmpty | NotEmpty rule control
func NotEmpty(input string) bool {
	return strings.TrimSpace(input) != ""
}
