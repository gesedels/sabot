// Package note implements the Note type and methods.
package note

import (
	"fmt"
	"time"

	"github.com/gesedels/sabot/sabot/tools/bolt"
	"github.com/gesedels/sabot/sabot/tools/neat"
	"go.etcd.io/bbolt"
)

// Note is a single recorded Note in a database.
type Note struct {
	DB   *bbolt.DB
	Name string
}

// New returns a new Note.
func New(db *bbolt.DB, name string) *Note {
	return &Note{db, name}
}

// Body returns the Note's body string.
func (n *Note) Body() (string, error) {
	body, err := bolt.GetValue(n.DB, n.Name, "body")
	if err != nil {
		return "", fmt.Errorf("cannot read Note %q - %w", n.Name, err)
	}

	return neat.Body(body), nil
}

// Check returns true if the Note's stored hash matches its body.
func (n *Note) Check() (bool, error) {
	body, err := bolt.GetValue(n.DB, n.Name, "body")
	if err != nil {
		return false, fmt.Errorf("cannot read Note %q - %w", n.Name, err)
	}

	hash, err := bolt.GetValue(n.DB, n.Name, "hash")
	if err != nil {
		return false, fmt.Errorf("cannot read Note %q - %w", n.Name, err)
	}

	body = neat.Body(body)
	return neat.Hash(body) == hash, nil
}

// Delete deletes the Note.
func (n *Note) Delete() error {
	if err := bolt.Delete(n.DB, n.Name); err != nil {
		return fmt.Errorf("cannot delete Note %q - %w", n.Name, err)
	}

	return nil
}

// Exists returns true if the Note exists.
func (n *Note) Exists() (bool, error) {
	ok, err := bolt.Exists(n.DB, n.Name)
	if err != nil {
		return false, fmt.Errorf("cannot read Note %q - %w", n.Name, err)
	}

	return ok, nil
}

// Init returns the Note's creation time.
func (n *Note) Init() (time.Time, error) {
	stmp, err := bolt.GetValue(n.DB, n.Name, "init")
	if err != nil {
		return time.Unix(0, 0), fmt.Errorf("cannot read Note %q - %w", n.Name, err)
	}

	return neat.Time(stmp), nil
}

// Last returns the Note's last edit time.
func (n *Note) Last() (time.Time, error) {
	stmp, err := bolt.GetValue(n.DB, n.Name, "last")
	if err != nil {
		return time.Unix(0, 0), fmt.Errorf("cannot read Note %q - %w", n.Name, err)
	}

	return neat.Time(stmp), nil
}

// Update overwrites the Note's body with a string.
func (n *Note) Update(body string) error {
	body = neat.Body(body)
	pairs := map[string]string{
		"body": body,
		"hash": neat.Hash(body),
		"last": neat.Stamp(time.Now()),
	}

	if err := bolt.Set(n.DB, n.Name, pairs); err != nil {
		return fmt.Errorf("cannot update Note %q - %w", n.Name, err)
	}

	return nil
}
