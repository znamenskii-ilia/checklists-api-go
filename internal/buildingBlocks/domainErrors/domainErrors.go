package domainErrors

import "errors"

var (
	ErrEntityConflict = errors.New("entity conflict")
	ErrEntityNotFound = errors.New("entity not found")
)
