package main

import "net/http"

// UserService provides the user management handlers
type UserService struct {
}

// NewUserService build a UserService.
func NewUserService() *UserService {
	return &UserService{}
}

// NewHandler handles user creation request
func (us *UserService) NewHandler(r *http.Request) (interface{}, error) {

	return responseMessage("created user"), nil
}
