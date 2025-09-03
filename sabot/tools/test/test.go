// Package test implements unit testing data and functions.
package test

import (
	"testing"

	"github.com/gesedels/sabot/sabot/tools/sqls"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// MockData is a string of database insert statements for unit testing.
const MockData = `
	insert into Flags (name, flag) values ('is-testing', true);
	insert into Notes (name) values ('alpha');
	insert into Notes (name) values ('bravo');
	insert into Pages (note, body, hash) values (1, 'Alpha one.', 'jZbLHn9hjSdEa6NSaLTltCsWOhrbqMv2tKN0d8zJT0w');
	insert into Pages (note, body, hash) values (1, 'Alpha two.', 'UY8K3s_jSS6gFYZUYOIj9K0XvOepidcQo8dJHfhFAqQ');
	insert into Pages (note, body, hash) values (2, 'Bravo one.', '<invalid hash>');
`

// AssertSQL asserts the first result of an SQL query.
func AssertSQL(t *testing.T, db *sqlx.DB, code string, want any) {
	var data any
	err := db.Get(&data, code)
	assert.NoError(t, err)

	switch want := want.(type) {
	case int:
		assert.Equal(t, int64(want), data)
	default:
		assert.Equal(t, want, data)
	}
}

// MockDB returns an in-memory database with default pragma, schema and mock data.
func MockDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec(sqls.Pragma + sqls.Schema + MockData); err != nil {
		t.Fatal(err)
	}

	return db
}
