// Package test implements unit testing data and functions.
package test

import (
	"path/filepath"
	"testing"

	"github.com/gesedels/sabot/sabot/tools/sqls"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// MockNotes is a slice of Notes table inserts for unit testing.
var MockNotes = [][]any{
	{1000, "alpha"},
	{2000, "bravo"},
	{3000, "empty"},
}

// MockPages is a slice of Pages table inserts for unit testing.
var MockPages = [][]any{
	{1000, 1, "Alpha one."},
	{1100, 1, "Alpha two."},
	{2000, 2, "Bravo one."},
}

// GetBool returns a database query as a boolean.
func GetBool(db *sqlx.DB, code string, elems ...any) bool {
	var data bool
	if err := db.Get(&data, code, elems...); err != nil {
		panic(err)
	}

	return data
}

// GetInt returns a database query as an integer.
func GetInt(db *sqlx.DB, code string, elems ...any) int {
	var data int
	if err := db.Get(&data, code, elems...); err != nil {
		panic(err)
	}

	return data
}

// GetMap returns a database query as a string:any map.
func GetMap(db *sqlx.DB, code string, elems ...any) map[string]any {
	data := make(map[string]any)
	rows, err := db.Queryx(code, elems...)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		if err := rows.MapScan(data); err != nil {
			panic(err)
		}
	}

	return data
}

// MockDB returns an in-memory database populated with mock data.
func MockDB(t *testing.T) *sqlx.DB {
	dire := t.TempDir()
	dest := filepath.Join(dire, "test.db")
	db := sqlx.MustConnect("sqlite3", dest)
	db.MustExec(sqls.Pragma + sqls.Schema)

	for _, note := range MockNotes {
		db.MustExec(
			"insert into Notes (init, name) values (?, ?)",
			note[0], note[1],
		)
	}

	for _, page := range MockPages {
		db.MustExec(
			"insert into Pages (init, note, body) values (?, ?, ?)",
			page[0], page[1], page[2],
		)
	}

	return db
}
