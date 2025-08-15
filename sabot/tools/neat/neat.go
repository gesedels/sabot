// Package neat implements value conversion and sanitisation functions.
package neat

import (
	"strings"
	"time"
	"unicode"
)

// Body returns a whitespace-trimmed body string.
func Body(body string) string {
	return strings.TrimSpace(body)
}

// Name returns a lowercase alphanumeric-with-dashes name string.
func Name(name string) string {
	var chars []rune
	for _, char := range strings.ToLower(name) {
		switch {
		case unicode.IsLetter(char) || unicode.IsNumber(char):
			chars = append(chars, char)
		case char == ' ' || char == '-' || char == '_':
			chars = append(chars, '-')
		}
	}

	return strings.Trim(string(chars), "-")
}

// Time returns a local Time object from a Unix UTC integer.
func Time(unix int64) time.Time {
	return time.Unix(unix, 0).In(time.Local)
}
