// Package page implements the Page type and methods.
package page

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Page is a single recorded version of a Note.
type Page struct {
	DB   *sqlx.DB `db:"-"`
	ID   int      `db:"id"`
	Init int      `db:"init"`
	Note int      `db:"note"`
	Body string   `db:"body"`
}

// Get returns an existing Page by ID.
func Get(db *sqlx.DB, id int) (*Page, error) {
	page := &Page{DB: db}
	code := "select * from Pages where id=? limit 1"
	err := db.Get(page, code, id)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot read Page %d - %w", id, err)
	default:
		return page, nil
	}
}

// Delete deletes the Page if it exists.
func (p *Page) Delete() error {
	code := "delete from Pages where id=?"
	if _, err := p.DB.Exec(code, p.ID); err != nil {
		return fmt.Errorf("cannot delete Page %d - %w", p.ID, err)
	}

	return nil
}
