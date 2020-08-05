package functions

import (
	"net"
)

//IsIP |
func IsIP(input string) bool {
	result := net.ParseIP(input)
	return result != nil
}
