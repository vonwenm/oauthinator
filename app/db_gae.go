// +build appengine

package oauthinator

import (
	"fmt"

	"appengine"
	"appengine/datastore"
)

type GaeUserStore struct {
}

// Key generate a key for the current user
func (u *User) Key(c appengine.Context) *datastore.Key {

	if u.Login == "" {
		panic("Tried Key on User with empty Login")
	}

	return datastore.NewKey(c, "User", u.Login, 0, nil)
}

func (*GaeUserStore) PutUser(c appengine.Context, u *User) error {

	if err := u.Valid(); err != nil {
		return fmt.Errorf("putting Commit: %v", err)
	}

	if _, err := datastore.Put(c, u.Key(c), u); err != nil {
		return fmt.Errorf("putting Commit: %v", err)
	}

	return nil
}
