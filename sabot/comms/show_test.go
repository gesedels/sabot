package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowCommand(t *testing.T) {
	// setup
	setupCommand(t, NewShowCommand)

	// success
	text, _, err := runCommand(t, "show", "alpha")
	assert.Equal(t, "Alpha.\n", text)
	assert.NoError(t, err)
}
