// Package comm implements the Command interface and collections.
package comm

import (
	"fmt"
	"io"

	"github.com/gesedels/sabot/sabot/items/book"
	"github.com/gesedels/sabot/sabot/tools/clui"
)

// Command is a callable user-facing command.
type Command interface {
	// Name returns the Command's callable name.
	Name() string

	// Help returns the Command's demo and help strings.
	Help() (string, string)

	// Run executes the Command with arguments.
	Run(io.Writer, []string) error
}

// NewCommand is a function that returns a new initialised Command.
type NewCommand func(*book.Book) Command

// Commands is a map of all defined Commands.
var Commands = map[string]NewCommand{
	"show": NewShowCommand,
}

// Run discovers and executes a Command with arguments.
func Run(w io.Writer, book *book.Book, elems []string) error {
	name, elems := clui.Split(elems)
	cfun, ok := Commands[name]
	if !ok {
		return fmt.Errorf("cannot run Command %q - does not exist", name)
	}

	return cfun(book).Run(w, elems)
}
