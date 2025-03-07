package iterations

import (
	"strings"
)

// The function takes in string and repeats the string five times
func Repeated(input string) string{
	var sb strings.Builder
	for i := 0; i < 5; i++ {
		sb.WriteString(input)
	}
	return sb.String()
}