package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("这是一个错误")
	fmt.Println(err.Error()) // 这是一个错误
}
