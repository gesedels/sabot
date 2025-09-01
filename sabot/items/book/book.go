// Package book implements the Book type and methods.
package book

import (
	"fmt"
	"time"

	"github.com/gesedels/sabot/sabot/items/note"
	"github.com/gesedels/sabot/sabot/tools/bolt"
	"github.com/gesedels/sabot/sabot/tools/neat"
	"go.etcd.io/bbolt"
)

// Book is a database containing multiple Notes.
type Book struct {
	DB *bbolt.DB
}

// toNotes returns a Note slice from a name string slice.
func toNotes(db *bbolt.DB, names []string) []*note.Note {
	var notes []*note.Note
	for _, name := range names {
		notes = append(notes, note.New(db, name))
	}

	return notes
}

// New returns a new Book.
func New(db *bbolt.DB) *Book {
	return &Book{db}
}

// NewPath returns a new Book from a database path.
func NewPath(path string) (*Book, error) {
	db, err := bbolt.Open(path, 0660, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("cannot open Book %q - %w", path, err)
	}

	return &Book{db}, nil
}

// Check returns all Notes failing a hash check.
func (b *Book) Check() ([]*note.Note, error) {
	names, err := bolt.List(b.DB)
	if err != nil {
		return nil, fmt.Errorf("cannot read Book - %w", err)
	}

	var notes []*note.Note
	for _, note := range toNotes(b.DB, names) {
		ok, err := note.Check()
		switch {
		case err != nil:
			return nil, fmt.Errorf("cannot read Book - %w", err)
		case !ok:
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// Close closes the Book's database.
func (b *Book) Close() error {
	if err := b.DB.Close(); err != nil {
		return fmt.Errorf("cannot close Book - %w", err)
	}

	return nil
}

// Create creates and returns a new Note.
func (b *Book) Create(name, body string) (*note.Note, error) {
	name = neat.Name(name)
	ok, err := bolt.Exists(b.DB, name)
	switch {
	case err != nil:
		return nil, fmt.Errorf("cannot create Note %q - %w", name, err)
	case ok:
		return nil, fmt.Errorf("cannot create Note %q - already exists", name)
	}

	body = neat.Body(body)
	pairs := map[string]string{
		"body": body,
		"hash": neat.Hash(body),
		"init": neat.Stamp(time.Now()),
		"last": neat.Stamp(time.Now()),
	}

	if err := bolt.Set(b.DB, name, pairs); err != nil {
		return nil, fmt.Errorf("cannot create Note %q - %w", name, err)
	}

	return note.New(b.DB, name), nil
}

// Get returns an existing Note.
func (b *Book) Get(name string) (*note.Note, error) {
	name = neat.Name(name)
	ok, err := bolt.Exists(b.DB, name)
	switch {
	case err != nil:
		return nil, fmt.Errorf("cannot read Note %q - %w", name, err)
	case !ok:
		return nil, fmt.Errorf("cannot read Note %q - does not exist", name)
	}

	return note.New(b.DB, name), nil
}

// GetOrCreate returns a newly created or existing Note.
func (b *Book) GetOrCreate(name string) (*note.Note, error) {
	name = neat.Name(name)
	ok, err := bolt.Exists(b.DB, name)
	switch {
	case err != nil:
		return nil, fmt.Errorf("cannot read Note %q - %w", name, err)
	case !ok:
		return b.Create(name, "")
	default:
		return b.Get(name)
	}
}

// List returns all existing Notes.
func (b *Book) List() ([]*note.Note, error) {
	names, err := bolt.List(b.DB)
	if err != nil {
		return nil, fmt.Errorf("cannot read Book - %w", err)
	}

	return toNotes(b.DB, names), nil
}

// Match returns all existing Notes with names containing a substring.
func (b *Book) Match(subs string) ([]*note.Note, error) {
	names, err := bolt.Match(b.DB, subs)
	if err != nil {
		return nil, fmt.Errorf("cannot read Book - %w", err)
	}

	return toNotes(b.DB, names), nil
}

// Search returns all existing Notes with bodies containing a substring.
func (b *Book) Search(subs string) ([]*note.Note, error) {
	names, err := bolt.Search(b.DB, "body", subs)
	if err != nil {
		return nil, fmt.Errorf("cannot read Book - %w", err)
	}

	return toNotes(b.DB, names), nil
}
