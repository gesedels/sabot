package book

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/gesedels/sabot/sabot/tools/bolt"
	"github.com/gesedels/sabot/sabot/tools/neat"
	"github.com/gesedels/sabot/sabot/tools/test"
	"github.com/stretchr/testify/assert"
)

func mockBook(t *testing.T) *Book {
	db := test.MockDB(t)
	return New(db)
}

func TestToNotes(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	notes := toNotes(db, []string{"alpha"})
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name)
}

func TestNew(t *testing.T) {
	// success
	book := mockBook(t)
	assert.NotNil(t, book.DB)
}

func TestNewPath(t *testing.T) {
	// setup
	path := filepath.Join(t.TempDir(), "bolt.db")

	// success
	book, err := NewPath(path)
	assert.NotNil(t, book.DB)
	assert.NoError(t, err)

	// error - database error
	book, err = NewPath("")
	assert.Nil(t, book)
	assert.EqualError(t, err, `cannot open Book "" - open : no such file or directory`)
}

func TestCheck(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes, err := book.Check()
	assert.Len(t, notes, 1)
	assert.Equal(t, "bravo", notes[0].Name)
	assert.NoError(t, err)

	// setup
	book.DB.Close()

	// error - database error
	notes, err = book.Check()
	assert.Nil(t, notes)
	assert.EqualError(t, err, `cannot read Book - database not open`)
}

func TestClose(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	err := book.Close()
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	note, err := book.Create("name", "Body.\n")
	assert.Equal(t, "name", note.Name)
	assert.NoError(t, err)

	// confirm
	pairs, _ := bolt.Get(note.DB, "name")
	assert.Equal(t, map[string]string{
		"body": "Body.\n",
		"hash": neat.Hash("Body.\n"),
		"init": time.Now().Format(time.RFC3339),
		"last": time.Now().Format(time.RFC3339),
	}, pairs)

	// error - already exists
	note, err = book.Create("name", "Body.\n")
	assert.Nil(t, note)
	assert.EqualError(t, err, `cannot create Note "name" - already exists`)

	// setup
	book.DB.Close()

	// error - database error
	note, err = book.Create("name", "Body.\n")
	assert.Nil(t, note)
	assert.EqualError(t, err, `cannot create Note "name" - database not open`)
}

func TestGet(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	note, err := book.Get("alpha")
	assert.Equal(t, "alpha", note.Name)
	assert.NoError(t, err)

	// error - does not exist
	note, err = book.Get("nope")
	assert.Nil(t, note)
	assert.EqualError(t, err, `cannot read Note "nope" - does not exist`)

	// setup
	book.DB.Close()

	// error - database error
	note, err = book.Get("alpha")
	assert.Nil(t, note)
	assert.EqualError(t, err, `cannot read Note "alpha" - database not open`)
}

func TestGetOrCreate(t *testing.T) {
	// setup
	book := mockBook(t)

	// success - create Note
	note, err := book.GetOrCreate("name")
	assert.Equal(t, "name", note.Name)
	assert.NoError(t, err)

	// confirm
	pairs, _ := bolt.Get(note.DB, "name")
	assert.Equal(t, map[string]string{
		"body": "\n",
		"hash": neat.Hash("\n"),
		"init": time.Now().Format(time.RFC3339),
		"last": time.Now().Format(time.RFC3339),
	}, pairs)

	// success - get Note
	note, err = book.GetOrCreate("name")
	assert.Equal(t, "name", note.Name)
	assert.NoError(t, err)

	// setup
	book.DB.Close()

	// error - database error
	note, err = book.GetOrCreate("name")
	assert.Nil(t, note)
	assert.EqualError(t, err, `cannot read Note "name" - database not open`)
}

func TestList(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes, err := book.List()
	assert.Len(t, notes, 2)
	assert.Equal(t, "alpha", notes[0].Name)
	assert.Equal(t, "bravo", notes[1].Name)
	assert.NoError(t, err)

	// setup
	book.DB.Close()

	// error - database error
	notes, err = book.List()
	assert.Nil(t, notes)
	assert.EqualError(t, err, `cannot read Book - database not open`)
}

func TestMatch(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes, err := book.Match("ALPH")
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name)
	assert.NoError(t, err)

	// setup
	book.DB.Close()

	// error - database error
	notes, err = book.Match("")
	assert.Nil(t, notes)
	assert.EqualError(t, err, `cannot read Book - database not open`)
}

func TestSearch(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes, err := book.Search("ALPH")
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name)
	assert.NoError(t, err)

	// setup
	book.DB.Close()

	// error - database error
	notes, err = book.Search("")
	assert.Nil(t, notes)
	assert.EqualError(t, err, `cannot read Book - database not open`)
}
