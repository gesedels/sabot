package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.", body)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME 123!!!\n")
	assert.Equal(t, "name-123", name)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Unix(1000, 0).In(time.Local)

	// success
	tobj := Time(1000)
	assert.Equal(t, want, tobj)
}
