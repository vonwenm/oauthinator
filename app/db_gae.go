// +build appengine

package main

import (
	"fmt"

	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
)

// Key generate a user key
func (u *User) Key(c context.Context) *datastore.Key {

	if u.Login == "" {
		panic("Tried Key on User with empty Login")
	}

	return datastore.NewKey(c, "User", u.Login, 0, nil)
}

func putUser(c context.Context, u *User) error {

	if err := u.Valid(); err != nil {
		return fmt.Errorf("putting Commit: %v", err)
	}

	if _, err := datastore.Put(c, u.Key(c), u); err != nil {
		return fmt.Errorf("putting Commit: %v", err)
	}

	return nil
}
