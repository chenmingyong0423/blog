package main

import (
	"errors"
	"fmt"
)

type UserNotError struct {
	Name string
}

func (e *UserNotError) Error() string {
	return fmt.Sprintf("user not found: name=%s", e.Name)
}

func main() {
	var err = &UserNotError{Name: "陈明勇"}
	var errUserNotFound = &UserNotError{}
	if errors.As(err, &errUserNotFound) {
		fmt.Println(errUserNotFound.Name) // 陈明勇
	} else {
		fmt.Println(err)
	}
}
