package utils

import (
	"fmt"
	"strings"
)

func ConcatStrings(s ...string) (string, error) {
	var builder strings.Builder

	for _, str := range s {
		_, err := builder.WriteString(str)
		if err != nil {
			return "", fmt.Errorf("error concat string %s", str)
		}
	}

	str := builder.String()

	builder.Reset()

	return str, nil
}
