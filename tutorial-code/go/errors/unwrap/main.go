package main

import (
	"errors"
	"fmt"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

func main() {
	err := fmt.Errorf("%w: name=%s", ErrUserNotFound, "陈明勇")
	fmt.Println(err)
	fmt.Println(errors.Unwrap(err))
}
