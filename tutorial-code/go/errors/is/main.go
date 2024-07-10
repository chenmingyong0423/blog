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
	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("error is ErrUserNotFound")
	} else {
		fmt.Println(err)
	}
}
