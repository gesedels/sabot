package note

import (
	"testing"

	"github.com/gesedels/sabot/sabot/tools/test"
	"github.com/stretchr/testify/assert"
)

func xNote(name string) *Note {
	db := test.MockDB()
	note, _ := Get(db, name)
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
	assert.NoError(t, err)

	// failure - non-existent Page
	note, err = Get(db, "")
	assert.Nil(t, note)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	note := xNote("alpha")

	// success
	err := note.Delete()
	assert.NoError(t, err)

	// success - check database
	ok := test.GetBool(note.DB, "select exists (select 1 from Notes where id=1)")
	assert.False(t, ok)
}

func TestLatest(t *testing.T) {
	// setup
	note := xNote("alpha")

	// success
	page, err := note.Latest()
	assert.Equal(t, "Alpha two.", page.Body)
	assert.NoError(t, err)

	// setup
	note = xNote("charlie")

	// failure - non-existent Page
	page, err = note.Latest()
	assert.Nil(t, page)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	note := xNote("alpha")

	// success
	page, err := note.Update("Body.")
	assert.Equal(t, 4, page.ID)
	assert.Equal(t, "Body.", page.Body)
	assert.NoError(t, err)
}
