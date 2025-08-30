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
	hash := Hash("Body.\n")
	assert.Equal(t, "RCYc4kLhuZ1Sx9Kky228taSrUHvtm5swMGKpafr-HR4", hash)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME_123!\n")
	assert.Equal(t, "name-123", name)
}

func TestStamp(t *testing.T) {
	// setup
	tobj := time.Date(2000, time.January, 1, 12, 0, 0, 0, time.UTC)

	// success
	stmp := Stamp(tobj)
	assert.Equal(t, "2000-01-01T12:00:00Z", stmp)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Date(2000, time.January, 1, 12, 0, 0, 0, time.UTC)

	// success
	tobj := Time("2000-01-01T12:00:00Z")
	assert.Equal(t, want, tobj)
}
