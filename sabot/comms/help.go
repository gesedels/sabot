package comm

import (
	"fmt"
	"io"

	"github.com/gesedels/sabot/sabot/items/book"
)

// HelpCommand is a command for printing help messages.
type HelpCommand struct {
	Book *book.Book
}

// NewHelpCommand returns a new HelpCommand.
func NewHelpCommand(book *book.Book) Command {
	return &HelpCommand{book}
}

// Name returns the HelpCommand's callable name.
func (c *HelpCommand) Name() string {
	return "help"
}

// Help returns the HelpCommand's help and demo strings.
func (c *HelpCommand) Help() (string, string) {
	return "help", "Print Sabot help."
}

// Run executes the HelpCommand with arguments.
func (c *HelpCommand) Run(w io.Writer, _ []string) error {
	fmt.Fprintln(w, "Available commands:")
	fmt.Fprintln(w)

	for name, newCommandFunc := range Commands {
		cmd := newCommandFunc(c.Book)
		demo, help := cmd.Help()
		fmt.Fprintf(w, "· %s: %s\n", name, help)
		fmt.Fprintf(w, "%-7s $ sabot %s\n", "", demo)
		fmt.Fprintln(w)
	}

	return nil
}
