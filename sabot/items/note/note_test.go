package note

import (
	"testing"

	"github.com/gesedels/sabot/sabot/tools/neat"
	"github.com/gesedels/sabot/sabot/tools/test"
	"github.com/stretchr/testify/assert"
)

func xNote() *Note {
	db := test.MockDB()
	note, _ := Get(db, "alpha")
	return note
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB()

	// success
	note, err := Get(db, "alpha")
	assert.NotNil(t, note.DB)
	assert.Equal(t, 1, note.ID)
	assert.Equal(t, 1000, note.Init)
	assert.Equal(t, "alpha", note.Name)
	assert.Equal(t, neat.Hash("alpha"), note.Hash)
	assert.NoError(t, err)

	// failure - non-existent Page
	note, err = Get(db, "nope")
	assert.Nil(t, note)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	note := xNote()

	// success
	err := note.Delete()
	assert.NoError(t, err)

	// success - check database
	var ok bool
	note.DB.Get(&ok, "select exists (select 1 from Notes where name='alpha')")
	assert.False(t, ok)
}

func TestLatest(t *testing.T) {
	// setup
	note := xNote()

	// success
	page, err := note.Latest()
	assert.Equal(t, "Alpha two.", page.Body)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	note := xNote()

	// success
	page, err := note.Update("Body.")
	assert.Equal(t, 4, page.ID)
	assert.Equal(t, "Body.", page.Body)
	assert.NoError(t, err)

	// success - check database
	var ok bool
	page.DB.Get(&ok, "select exists (select 1 from Pages where id=4)")
	assert.True(t, ok)
}
