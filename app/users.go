package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// User which represents a user of the system
type User struct {
	Login   string `json:"login,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Bio     string `json:"bio,omitempty" datastore:",noindex"`
	URL     string `json:"url,omitempty"`
	Tz      string `json:"tz,omitempty"`
	Org     string `json:"org,omitempty"`
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

// UserService provides the user management handlers
type UserService struct {
}

// NewUserService build a UserService.
func NewUserService() *UserService {
	return &UserService{}
}

// NewHandler handles user creation request
func (us *UserService) NewHandler(r *http.Request) (interface{}, error) {

	c := newContext(r)

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		errorf(c, "failed to parse user: %v", err)
		return nil, ErrInvalidForm
	}

	user.Created = time.Now()

	err = putUser(c, &user)

	if err != nil {
		errorf(c, "failed to save user: %v", err)
		return nil, ErrInvalidForm
	}

	return user, nil
}
