// Package test implements unit testing data and functions.
package test

import (
	"github.com/gesedels/sabot/sabot/tools/sqls"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// MockNotes is a slice of Notes table inserts for unit testing.
var MockNotes = [][]any{
	{1000, "alpha"},
	{2000, "bravo"},
	{3000, "charlie"},
}

// MockPages is a slice of Pages table inserts for unit testing.
var MockPages = [][]any{
	{1000, 1, "Alpha one."},
	{1100, 1, "Alpha two."},
	{2000, 2, "Bravo one."},
}

// GetBool returns a database query as a boolean.
func GetBool(db *sqlx.DB, code string, elems ...any) bool {
	var ok bool
	if err := db.Get(&ok, code, elems...); err != nil {
		panic(err)
	}

	return ok
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
func MockDB() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")
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
