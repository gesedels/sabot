package comm

import (
	"fmt"
	"io"

	"github.com/gesedels/sabot/sabot/items/book"
)

// ShowCommand is a command for printing existing Notes.
type ShowCommand struct {
	Book *book.Book
}

// NewShowCommand returns a new ShowCommand.
func NewShowCommand(book *book.Book) Command {
	return &ShowCommand{book}
}

// Name returns the ShowCommand's callable name.
func (c *ShowCommand) Name() string {
	return "show"
}

// Help returns the ShowCommand's help and demo strings.
func (c *ShowCommand) Help() (string, string) {
	return "show NOTES...", "Print an existing note"
}

// Run executes the ShowCommand with arguments.
func (c *ShowCommand) Run(w io.Writer, elems []string) error {
	for _, name := range elems {
		note, err := c.Book.Get(name)
		if err != nil {
			return err
		}

		body, err := note.Body()
		if err != nil {
			return err
		}

		fmt.Fprint(w, body)
	}

	return nil
}
