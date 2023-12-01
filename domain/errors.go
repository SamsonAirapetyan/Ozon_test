package domain

import "errors"

var (
	ErrNoRecordFound   = errors.New("no record has been found")
	ErrInvalidArgument = errors.New("invalid argument")
)
