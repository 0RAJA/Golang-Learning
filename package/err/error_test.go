package err

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewErr(t *testing.T) {
	err := NewErr(TimeOutCode, errors.New("222"))
	err2 := NewErr(TimeOutCode, err)
	fmt.Println(errors.Unwrap(err2))
	fmt.Println(errors.Unwrap(err))
}
