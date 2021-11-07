package domain

import "errors"

var (
	ErrInternalServer = errors.New("internal Server Error")
	ErrNotFound = errors.New("your requested Item is not found")
	ErrConflict = errors.New("your Item already exist")
	ErrBadParamInput = errors.New("given Param is not valid")
)