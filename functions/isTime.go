package functions

import "time"

//IsTime rule
func IsTime(input string, format string) bool {

	_, err := time.Parse(format, input)

	return err == nil
	
}
