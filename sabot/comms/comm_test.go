package comm

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/gesedels/sabot/sabot/items/book"
	"github.com/gesedels/sabot/sabot/tools/test"
	"github.com/stretchr/testify/assert"
)

type mockCommand struct {
	Book *book.Book
}

func newMockCommand(book *book.Book) Command  { return &mockCommand{book} }
func (c *mockCommand) Name() string           { return "mock" }
func (c *mockCommand) Help() (string, string) { return "demo", "help" }
func (c *mockCommand) Run(w io.Writer, elems []string) error {
	fmt.Fprintf(w, "elems=%v", elems)
	return nil
}

func setupCommand(t *testing.T, cfun NewCommand) {
	comm := cfun(nil)
	name := comm.Name()
	demo, help := comm.Help()
	assert.NotEmpty(t, name)
	assert.NotEmpty(t, demo)
	assert.NotEmpty(t, help)
	Commands[name] = cfun
}

func runCommand(t *testing.T, elems ...string) (string, *book.Book, error) {
	b := new(bytes.Buffer)
	book := book.New(test.DB(t))
	err := Run(b, book, elems)
	return b.String(), book, err
}

func TestRun(t *testing.T) {
	// setup
	b := new(bytes.Buffer)
	Commands["mock"] = newMockCommand

	// success
	err := Run(b, nil, []string{"mock", "argument"})
	assert.Equal(t, "elems=[argument]", b.String())
	assert.NoError(t, err)

	// error - does not exist
	err = Run(nil, nil, []string{"nope"})
	assert.EqualError(t, err, `cannot run Command "nope" - does not exist`)
}
