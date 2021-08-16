package main

import (
	"errors"
	"fmt"
	"os"
)

type ListErr struct {
	msg string
	err error
}

func (e *ListErr) Error() string {
	return e.msg
}

func (e *ListErr) Unwrap() error {
	return e.err
}

func main() {
	err1 := fmt.Errorf("%w", os.ErrPermission)
	err2 := fmt.Errorf("%w", err1)
	fmt.Println(errors.Is(err2, err1))
}
