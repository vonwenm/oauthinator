// +build !appengine

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/boltdb/bolt"
	"golang.org/x/net/context"
)

var (
	defaultStore = newStore()
)

type store struct {
	db *bolt.DB
}

func (s *store) save(entity string, key string, val interface{}) error {

	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(val)
	if err != nil {
		return err
	}

	return s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(entity)).Put([]byte(key), buf.Bytes())
	})
}

func newStore() *store {
	db, err := bolt.Open("./oauthinator.db", 0666, nil)

	if err != nil {
		panic(fmt.Sprintf("unable to create data store: %v", err))
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("Users"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		logf(nil, "created bucket %v", b.Stats())

		return nil
	})

	if err != nil {
		panic(fmt.Sprintf("unable to create data store: %v", err))
	}

	return &store{db}
}

// Key generate a user key
func (u *User) Key(_ context.Context) string {

	if u.Login == "" {
		panic("Tried Key on User with empty Login")
	}

	return fmt.Sprintf("User-%v", u.Login)
}

func putUser(_ context.Context, u *User) error {
	if u == nil {
		return fmt.Errorf("User is nil")
	}

	if err := u.Valid(); err != nil {
		return fmt.Errorf("putting User: %v", err)
	}

	return defaultStore.save("Users", u.Key(nil), u)
}
