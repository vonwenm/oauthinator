package oauthinator

import (
	"errors"
	"time"
)

// User which represents a user of the system
type User struct {
	Login   string
	Name    string
	Email   string
	Bio     string `datastore:",noindex"`
	URL     string
	Tz      string
	Org     string
	Created time.Time
}

// Valid some simple validation of the user used to check creation and update
func (u *User) Valid() error {
	if u.Login == "" {
		return errors.New("invalid user, login required")
	}
	if u.Name == "" {
		return errors.New("invalid user, name required")
	}
	if u.Email == "" {
		return errors.New("invalid user, email required")
	}
	return nil
}
