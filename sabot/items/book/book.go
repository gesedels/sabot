// Package book implements the Book type and methods.
package book

import (
	"database/sql"
	"fmt"

	"github.com/gesedels/sabot/sabot/items/note"
	"github.com/gesedels/sabot/sabot/tools/neat"
	"github.com/gesedels/sabot/sabot/tools/sqls"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Book is a single initialised database with multiple notes.
type Book struct {
	DB *sqlx.DB `db:"-"`
}

// Open returns a new Book with an initialised database connection.
func Open(path string) (*Book, error) {
	db, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("cannot open database %q - %w", path, err)
	}

	if _, err := db.Exec(sqls.Pragma); err != nil {
		return nil, fmt.Errorf("cannot initialise database %q - %w", path, err)
	}

	var size int
	code := "select count(*) from SQLITE_SCHEMA"
	if err := db.Get(&size, code); err != nil {
		return nil, fmt.Errorf("cannot initialise database %q - %w", path, err)
	}

	if size == 0 {
		if _, err := db.Exec(sqls.Schema); err != nil {
			return nil, fmt.Errorf("cannot initialise database %q - %w", path, err)
		}
	}

	return &Book{db}, nil
}

// Create creates and returns a new Note with one Page.
func (b *Book) Create(name, body string) (*note.Note, error) {
	name = neat.Name(name)
	code := "insert into Notes (name) values (?)"
	if _, err := b.DB.Exec(code, name); err != nil {
		return nil, fmt.Errorf("cannot create Note %q - %w", name, err)
	}

	note, err := note.Get(b.DB, name)
	if err != nil {
		return nil, err
	}

	body = neat.Body(body)
	if _, err := note.Update(body); err != nil {
		return nil, err
	}

	return note, nil
}

// Get returns an existing Note.
func (b *Book) Get(name string) (*note.Note, error) {
	return note.Get(b.DB, name)
}

// Match returns all existing Notes with names containing a substring.
func (b *Book) Match(text string) ([]*note.Note, error) {
	text = "%" + text + "%"
	code := "select * from Notes where name like ? order by name asc"
	return b.Select(code, text)
}

// Select returns all Notes matching a select query.
func (b *Book) Select(code string, elems ...any) ([]*note.Note, error) {
	var notes []*note.Note
	err := b.DB.Select(&notes, code, elems...)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot read Notes - %w", err)
	default:
		for _, note := range notes {
			note.DB = b.DB
		}

		return notes, nil
	}
}
