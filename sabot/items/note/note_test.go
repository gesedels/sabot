package note

import (
	"testing"
	"time"

	"github.com/gesedels/sabot/sabot/tools/bolt"
	"github.com/gesedels/sabot/sabot/tools/neat"
	"github.com/gesedels/sabot/sabot/tools/test"
	"github.com/stretchr/testify/assert"
)

var (
	timeZero = time.Unix(0, 0)
	timeJan1 = time.Date(2000, time.January, 1, 12, 0, 0, 0, time.Local)
)

func mockNote(t *testing.T) *Note {
	db := test.DB(t)
	return New(db, "alpha")
}

func TestNew(t *testing.T) {
	// success
	note := mockNote(t)
	assert.NotNil(t, note.DB)
	assert.Equal(t, "alpha", note.Name)
}

func TestBody(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	body, err := note.Body()
	assert.Equal(t, "Alpha.\n", body)
	assert.NoError(t, err)

	// setup
	note.DB.Close()

	// error - database error
	body, err = note.Body()
	assert.Empty(t, body)
	assert.EqualError(t, err, `cannot read Note "alpha" - database not open`)
}

func TestCheck(t *testing.T) {
	// setup
	note := mockNote(t)

	// success - true
	ok, err := note.Check()
	assert.True(t, ok)
	assert.NoError(t, err)

	// setup
	bolt.Set(note.DB, "alpha", map[string]string{"body": ""})

	// success - false
	ok, err = note.Check()
	assert.False(t, ok)
	assert.NoError(t, err)

	// setup
	note.DB.Close()

	// error - database error
	ok, err = note.Check()
	assert.False(t, ok)
	assert.EqualError(t, err, `cannot read Note "alpha" - database not open`)
}

func TestDelete(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	err := note.Delete()
	assert.NoError(t, err)

	// confirm
	ok, _ := bolt.Exists(note.DB, "alpha")
	assert.False(t, ok)

	// setup
	note.DB.Close()

	// error - database error
	err = note.Delete()
	assert.EqualError(t, err, `cannot delete Note "alpha" - database not open`)
}

func TestExists(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	ok, err := note.Exists()
	assert.True(t, ok)
	assert.NoError(t, err)

	// setup
	note.DB.Close()

	// error - database error
	ok, err = note.Exists()
	assert.False(t, ok)
	assert.EqualError(t, err, `cannot read Note "alpha" - database not open`)
}

func TestInit(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	tobj, err := note.Init()
	assert.Equal(t, timeJan1, tobj)
	assert.NoError(t, err)

	// setup
	note.DB.Close()

	// error - database error
	tobj, err = note.Init()
	assert.Equal(t, timeZero, tobj)
	assert.EqualError(t, err, `cannot read Note "alpha" - database not open`)
}

func TestLast(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	tobj, err := note.Last()
	assert.Equal(t, timeJan1, tobj)
	assert.NoError(t, err)

	// setup
	note.DB.Close()

	// error - database error
	tobj, err = note.Last()
	assert.Equal(t, timeZero, tobj)
	assert.EqualError(t, err, `cannot read Note "alpha" - database not open`)
}

func TestUpdate(t *testing.T) {
	// setup
	note := mockNote(t)
	hash := neat.Hash("Body.\n")
	last := time.Now().Format(time.RFC3339)

	// success
	err := note.Update("Body.\n")
	assert.NoError(t, err)

	// confirm
	pairs, _ := bolt.Get(note.DB, "alpha")
	assert.Equal(t, "Body.\n", pairs["body"])
	assert.Equal(t, hash, pairs["hash"])
	assert.Equal(t, last, pairs["last"])

	// setup
	note.DB.Close()

	// error - database error
	err = note.Update("Body.\n")
	assert.EqualError(t, err, `cannot update Note "alpha" - database not open`)
}
