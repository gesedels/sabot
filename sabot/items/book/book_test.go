package book

import (
	"testing"

	"github.com/gesedels/sabot/sabot/tools/test"
	"github.com/stretchr/testify/assert"
)

func xBook(t *testing.T) *Book {
	db := test.MockDB(t)
	book := &Book{DB: db}
	return book
}

func TestOpen(t *testing.T) {
	// success
	book, err := Open(":memory:")
	assert.NotNil(t, book.DB)
	assert.NoError(t, err)

	// success - check database
	size := test.GetInt(book.DB, "select count(*) from SQLITE_SCHEMA")
	assert.NotZero(t, size)
}

func TestCreate(t *testing.T) {
	// setup
	book := xBook(t)

	// success
	note, err := book.Create("name", "Body.")
	assert.NoError(t, err)
	assert.Equal(t, 4, note.ID)
	assert.Equal(t, "name", note.Name)

	// success - check database
	data := test.GetMap(book.DB, "select * from Pages where note=4")
	assert.Equal(t, "Body.", data["body"])
}

func TestGet(t *testing.T) {
	// setup
	book := xBook(t)

	// success
	note, err := book.Get("alpha")
	assert.Equal(t, "alpha", note.Name)
	assert.NoError(t, err)
}

func TestMatch(t *testing.T) {
	// setup
	book := xBook(t)

	// success
	notes, err := book.Match("ALPH")
	assert.Len(t, notes, 1)
	// assert.Equal(t, "alpha", notes[0].Name)
	assert.NoError(t, err)
}

func TestSelect(t *testing.T) {
	// setup
	book := xBook(t)

	// success
	notes, err := book.Select("select * from Notes where name=?", "alpha")
	assert.Len(t, notes, 1)
	assert.NotNil(t, notes[0].DB)
	assert.Equal(t, "alpha", notes[0].Name)
	assert.NoError(t, err)

	// failure - no resulting Notes
	notes, err = book.Select("select * from Notes where name=?", "")
	assert.Nil(t, notes)
	assert.NoError(t, err)
}
