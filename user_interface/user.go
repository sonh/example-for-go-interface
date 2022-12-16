package user_interface

import (
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
	"net/mail"
)

type User interface {
	Email() string
}

type RandomUser struct{}

func (user RandomUser) Email() string {
	return ulid.Make().String() + "@example.com"
}

func ValidUser(user User) error {
	switch {
	case user.Email() == "":
		return errors.New("email can not be empty")
	case valid(user.Email()):
		return errors.New("email is not valid")
	}
	return nil
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
