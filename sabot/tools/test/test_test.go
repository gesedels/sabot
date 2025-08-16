package test

import (
	"path/filepath"
	"testing"

	"github.com/gesedels/sabot/sabot/tools/sqls"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetBool(t *testing.T) {
	// setup
	db := MockDB(t)

	// success
	ok := GetBool(db, "select 1")
	assert.True(t, ok)
}

func TestGetInt(t *testing.T) {
	// setup
	db := MockDB(t)

	// success
	data := GetInt(db, "select 1")
	assert.Equal(t, 1, data)
}

func TestGetMap(t *testing.T) {
	// setup
	db := MockDB(t)

	// success
	data := GetMap(db, "select * from Notes where id=1 limit 1")
	assert.Equal(t, "alpha", data["name"])
}

func TestMockDB(t *testing.T) {
	// success
	db := MockDB(t)
	assert.NotNil(t, db)
}

func TestMockInsert(t *testing.T) {
	// setup
	dire := t.TempDir()
	dest := filepath.Join(dire, "test.db")
	db := sqlx.MustConnect("sqlite3", dest)
	db.MustExec(sqls.Pragma + sqls.Schema)

	// success
	MockInsert(db)
	size := GetInt(db, "select count(*) from SQLITE_SCHEMA")
	assert.NotZero(t, size)
}
