package student_interface_embed_user_interface

import (
	"github.com/pkg/errors"

	"interace_homework/user_interface"
)

type Student interface {
	user_interface.User
	Grade() int
}

// ValidStudent func
// This student arg makes the func loose coupling with Student interface
func ValidStudent(student Student) error {
	if err := user_interface.ValidUser(student); err != nil {
		return err
	}

	switch {
	case student.Grade() < 0:
		return errors.New("student grade is not valid")
	}

	return nil
}
