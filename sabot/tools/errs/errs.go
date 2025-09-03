// Package errs implements error creation and formatting functions.
package errs

import (
	"fmt"
)

// Messages is a map of all defined error messages.
var Messages = map[string]string{
	"path-none": "cannot locate database path ($HOME not set)",
}

// Error returns a formatted error message from Errors.
func Error(name string, elems ...any) error {
	return fmt.Errorf(Messages[name], elems...)
}
