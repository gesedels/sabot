// Package clui implements command-line user interface functions.
package clui

import (
	"strings"
)

// Split returns a Command name and argument slice from an argument slice.
func Split(elems []string) (string, []string) {
	switch len(elems) {
	case 0:
		return "", nil
	case 1:
		return strings.ToLower(elems[0]), nil
	default:
		return strings.ToLower(elems[0]), elems[1:]
	}
}
