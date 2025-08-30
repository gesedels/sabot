// Package bolt implements Bolt database functions.
package bolt

import (
	"strings"

	"go.etcd.io/bbolt"
)

// Delete deletes an existing Bucket.
func Delete(db *bbolt.DB, name string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(name))
	})
}

// Exists returns true if a Bucket exists.
func Exists(db *bbolt.DB, name string) (bool, error) {
	var ok bool

	return ok, db.View(func(tx *bbolt.Tx) error {
		ok = tx.Bucket([]byte(name)) != nil
		return nil
	})
}

// Get returns an existing Bucket as a string map.
func Get(db *bbolt.DB, name string) (map[string]string, error) {
	var pairs map[string]string

	return pairs, db.View(func(tx *bbolt.Tx) error {
		if buck := tx.Bucket([]byte(name)); buck != nil {
			pairs = make(map[string]string)
			return buck.ForEach(func(attr []byte, data []byte) error {
				pairs[string(attr)] = string(data)
				return nil
			})
		}

		return nil
	})
}

// GetValue returns an existing Bucket value as a string.
func GetValue(db *bbolt.DB, name, attr string) (string, error) {
	var data string

	return data, db.View(func(tx *bbolt.Tx) error {
		if buck := tx.Bucket([]byte(name)); buck != nil {
			data = string(buck.Get([]byte(attr)))
		}

		return nil
	})
}

// List returns all existing Bucket names.
func List(db *bbolt.DB) ([]string, error) {
	var names []string

	return names, db.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
			names = append(names, string(name))
			return nil
		})
	})
}

// Match returns all existing Bucket names containing a substring.
func Match(db *bbolt.DB, text string) ([]string, error) {
	var goods []string

	names, err := List(db)
	text = strings.ToLower(text)
	for _, name := range names {
		if strings.Contains(strings.ToLower(name), text) {
			goods = append(goods, name)
		}
	}

	return goods, err
}

// Search returns all existing Bucket names with an attribute containing a substring.
func Search(db *bbolt.DB, attr, text string) ([]string, error) {
	var goods []string

	text = strings.ToLower(text)
	return goods, db.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(name []byte, buck *bbolt.Bucket) error {
			data := buck.Get([]byte(attr))
			if strings.Contains(strings.ToLower(string(data)), text) {
				goods = append(goods, string(name))
			}

			return nil
		})
	})
}

// Set overwrites a new or existing Bucket with a string map.
func Set(db *bbolt.DB, name string, pairs map[string]string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		buck, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}

		for attr, data := range pairs {
			if err := buck.Put([]byte(attr), []byte(data)); err != nil {
				return err
			}
		}

		return nil
	})
}
