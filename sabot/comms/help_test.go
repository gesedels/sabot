package comms

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const helpText = `
Available commands:

· help: Print Sabot help.
        $ sabot help

`

func TestHelpCommand(t *testing.T) {
	// setup
	Commands = map[string]NewCommand{"help": NewHelpCommand}
	setupCommand(t, NewHelpCommand)

	// success
	text, _, err := runCommand(t, "help")
	assert.Equal(t, strings.TrimLeft(helpText, "\n"), text)
	assert.NoError(t, err)
}
