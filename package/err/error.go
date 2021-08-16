package err

import (
	"fmt"
)

const (
	ErrorStrFormat = "Error Code : %v, Err : %v"
)

const (
	TimeOutCode = iota
)

type MyErr struct {
	Code int   `json:"code,omitempty"`
	Err  error `json:"err,omitempty"`
}

func (e *MyErr) Error() string {
	return fmt.Sprintf(ErrorStrFormat, e.Code, e.Err)
}

func (e *MyErr) Unwrap() error {
	return e.Err
}

func NewErr(code int, err error) error {
	return &MyErr{
		Code: code,
		Err:  err,
	}
}
