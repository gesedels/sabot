// Package test implements unit testing data and functions.
package test

import (
	"path/filepath"
	"testing"

	"github.com/gesedels/sabot/sabot/tools/neat"
	"go.etcd.io/bbolt"
)

// MockData is mock database data for unit testing.
var MockData = map[string]map[string]string{
	"alpha": {
		"body": "Alpha.\n",
		"hash": neat.Hash("Alpha.\n"),
		"init": "2000-01-01T12:00:00Z00:00",
		"last": "2000-01-01T12:00:00Z00:00",
	},

	"bravo": {
		"body": "Bravo #foo #bar.\n",
		"hash": neat.Hash("Bravo #foo #bar.\n"),
		"init": "2000-01-02T12:00:00Z00:00",
		"last": "2000-01-02T18:00:00Z00:00",
	},
}

// DB returns a temporary database populated with MockData.
func DB(t *testing.T) *bbolt.DB {
	path := filepath.Join(t.TempDir(), "bolt.db")
	db, _ := bbolt.Open(path, 0666, nil)
	db.Update(func(tx *bbolt.Tx) error {
		for name, pairs := range MockData {
			buck, _ := tx.CreateBucket([]byte(name))
			for attr, data := range pairs {
				buck.Put([]byte(attr), []byte(data))
			}
		}

		return nil
	})

	return db
}
