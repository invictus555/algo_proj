package utils

import "errors"

var (
	ErrOom   = errors.New("out of memory")
	ErrEmpty = errors.New("list is empty")
)
