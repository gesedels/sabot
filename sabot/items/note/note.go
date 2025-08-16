// Package note implements the Note type and methods.
package note

import (
	"database/sql"
	"fmt"

	"github.com/gesedels/sabot/sabot/items/page"
	"github.com/gesedels/sabot/sabot/tools/neat"
	"github.com/jmoiron/sqlx"
)

// Note is a single recorded note with multiple versions.
type Note struct {
	DB   *sqlx.DB `db:"-"`
	ID   int      `db:"id"`
	Init int      `db:"init"`
	Name string   `db:"name"`
}

// Get returns an existing Note by name.
func Get(db *sqlx.DB, name string) (*Note, error) {
	name = neat.Name(name)
	note := &Note{DB: db}
	code := "select * from Notes where name=? limit 1"
	err := db.Get(note, code, name)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot read Note %q - %w", name, err)
	default:
		return note, nil
	}
}

// Delete deletes the Note if it exists.
func (n *Note) Delete() error {
	code := "delete from Notes where id=?"
	if _, err := n.DB.Exec(code, n.ID); err != nil {
		return fmt.Errorf("cannot delete Note %q - %w", n.Name, err)
	}

	return nil
}

// Latest return the Note's latest Page.
func (n *Note) Latest() (*page.Page, error) {
	page := &page.Page{DB: n.DB}
	code := "select * from Pages where note=? order by id desc limit 1"
	err := n.DB.Get(page, code, n.ID)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot read Note %q - %w", n.Name, err)
	default:
		return page, nil
	}
}

// Update creates and returns a new Page in the Note.
func (n *Note) Update(body string) (*page.Page, error) {
	body = neat.Body(body)
	code := "insert into Pages (note, body) values (?, ?)"
	if _, err := n.DB.Exec(code, n.ID, body); err != nil {
		return nil, fmt.Errorf("cannot update Note %q - %w", n.Name, err)
	}

	return n.Latest()
}
