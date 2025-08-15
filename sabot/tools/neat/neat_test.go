package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.\n", body)
}

func TestHash(t *testing.T) {
	// success
	hash := Hash("Text.\n")
	assert.Equal(t, "YeuxYsUiVnuqvt-xQv5zQduD0q0nLyiGySwa7wR7RVE", hash)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME 123!!!\n")
	assert.Equal(t, "name-123", name)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Unix(1000, 0)
	zero := time.Unix(0, 0)

	// success - valid string
	tobj := Time("1000\n")
	assert.Equal(t, want, tobj)

	// success - invalid string
	tobj = Time("")
	assert.Equal(t, zero, tobj)
}
