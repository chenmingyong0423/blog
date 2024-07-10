package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err := errors.Join(err1, err2)
	fmt.Println(err)
	fmt.Println(errors.Is(err, err1)) // true
	fmt.Println(errors.Is(err, err2)) // true
}
