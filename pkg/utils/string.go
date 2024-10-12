package utils

import "strings"

var builder strings.Builder

func ConcatStrings(s []string) string {
	for _, string := range s {
		builder.WriteString(string)
	}

	str := builder.String()

	builder.Reset()

	return str
}
