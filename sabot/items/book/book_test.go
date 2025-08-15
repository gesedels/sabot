package book

import (
	"testing"

	"github.com/gesedels/sabot/sabot/items/page"
	"github.com/gesedels/sabot/sabot/tools/neat"
	"github.com/stretchr/testify/assert"
)

func xBook() *Book {
	book, _ := Open(":memory:")
	return book
}

func TestOpen(t *testing.T) {
	// success
	book, err := Open(":memory:")
	assert.NotNil(t, book.DB)
	assert.NoError(t, err)

	// success - check database
	var size int
	book.DB.Get(&size, "select count(*) from SQLITE_SCHEMA")
	assert.NotZero(t, size)
}

func TestCreate(t *testing.T) {
	// setup
	book := xBook()

	// success
	note, err := book.Create("name", "Body.")
	assert.NoError(t, err)
	assert.Equal(t, 1, note.ID)
	assert.Equal(t, "name", note.Name)

	// success - check database
	page := new(page.Page)
	book.DB.Get(page, "select * from Pages where note=1")
	assert.Equal(t, "Body.", page.Body)
	assert.Equal(t, neat.Hash("Body."), page.Hash)
}
