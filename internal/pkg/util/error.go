package util

import "errors"

type ErrorDefined error

var (
	ErrorWrongTypePageIdx  ErrorDefined = errors.New("page index wrong type")
	ErrorWrongTypePageSize ErrorDefined = errors.New("page size wrong type")
)
