// Package neat implements data sanitisation functions.
package neat

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"time"
	"unicode"
)

// Body returns a whitespace-trimmed body string with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Hash returns a base64-encoded SHA256 hash of a body string.
func Hash(body string) string {
	hash := sha256.Sum256([]byte(body))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

// Name returns a lowercase alphanumeric-with-dashes name string.
func Name(name string) string {
	var chars []rune
	for _, char := range strings.ToLower(name) {
		switch {
		case unicode.IsLetter(char) || unicode.IsNumber(char):
			chars = append(chars, char)
		case char == '-' || char == '_':
			chars = append(chars, '-')
		}
	}

	return string(chars)
}

// Stamp returns an RFC3339 timestamp string from a Time object.
func Stamp(tobj time.Time) string {
	return tobj.Format(time.RFC3339)
}

// Time returns a Time object from an RFC3339 timestamp string.
func Time(stmp string) time.Time {
	tobj, _ := time.Parse(time.RFC3339, stmp)
	return tobj
}
