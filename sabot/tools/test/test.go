// Package test implements unit testing data and functions.
package test

import (
	"github.com/gesedels/sabot/sabot/tools/neat"
	"github.com/gesedels/sabot/sabot/tools/sqls"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// MockNotes is a slice of Notes table inserts for unit testing.
var MockNotes = [][]string{
	{"1000", "alpha"},
	{"2000", "bravo"},
}

// MockPages is a slice of Pages table inserts for unit testing.
var MockPages = [][]string{
	{"1000", "1", "Alpha one.\n"},
	{"1100", "1", "Alpha two.\n"},
	{"2000", "2", "Bravo one.\n"},
}

// MockDB returns an in-memory database populated with mock data.
func MockDB() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	db.MustExec(sqls.Pragma + sqls.Schema)

	for _, note := range MockNotes {
		db.MustExec(
			"insert into Notes (init, name, hash) values (?, ?, ?)",
			note[0], note[1], neat.Hash(note[1]),
		)
	}

	for _, page := range MockPages {
		db.MustExec(
			"insert into Pages (init, note, body, hash) values (?, ?, ?, ?)",
			page[0], page[1], page[2], neat.Hash(page[2]),
		)
	}

	return db
}
