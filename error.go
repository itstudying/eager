package eager

import (
	"errors"
	"fmt"
)

// all errors
var (
	ErrNotPtr   = errors.New("config is not ptr")
	ErrNotExist = errors.New("file is not exist")
	ErrNotFile  = errors.New("is dir")
	ErrAssert   = errors.New("assert error")
)

func newError(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("error: %s , info: %s", msg, err)
	}
	return fmt.Errorf(msg)
}
