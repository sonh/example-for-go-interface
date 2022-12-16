package main

import (
	"fmt"
	"github.com/pkg/errors"

	"interace_homework/user_interface"
)

// Student also a User
type Student struct {
	user_interface.User
	Grade int
}

// ValidStudent func
// This student arg makes the func tight coupling with Student struct
func ValidStudent(student Student) error {
	if err := user_interface.ValidUser(student); err != nil {
		return err
	}

	switch {
	case student.Grade < 0:
		return errors.New("student grade is not valid")
	}

	return nil
}

// try to remove this comment to override method of User
/*func (student Student) Email() string {
	return "override_email"
}*/

func main() {
	student := Student{
		User:  user_interface.RandomUser{},
		Grade: 1,
	}

	fmt.Println("------------------------------------------------")
	fmt.Println("student email: ", student.Email())
}
