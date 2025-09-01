// Package clui implements command-line user interface functions.
package clui

import (
	"fmt"
	"strings"
)

// Need returns an error if an argument slice is below a minimum length.
func Need(elems []string, size int) error {
	if size > len(elems) {
		return fmt.Errorf("not enough arguments")
	}

	return nil
}

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
