package page

import (
	"testing"

	"github.com/gesedels/sabot/sabot/tools/test"
	"github.com/stretchr/testify/assert"
)

func xPage() *Page {
	db := test.MockDB()
	page, _ := Get(db, 1)
	return page
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB()

	// success
	page, err := Get(db, 1)
	assert.NotNil(t, page.DB)
	assert.Equal(t, 1, page.ID)
	assert.Equal(t, 1000, page.Init)
	assert.Equal(t, 1, page.Note)
	assert.Equal(t, "Alpha one.", page.Body)
	assert.NoError(t, err)

	// failure - non-existent Page
	page, err = Get(db, -1)
	assert.Nil(t, page)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	page := xPage()

	// success
	err := page.Delete()
	assert.NoError(t, err)

	// success - check database
	ok := test.GetBool(page.DB, "select exists (select 1 from Pages where id=1)")
	assert.False(t, ok)
}
