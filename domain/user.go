package domain

import (
	"errors"
)

type User struct {
	ID       string // UUID in string format
	Username string
	Email    string
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}
	return nil
}
