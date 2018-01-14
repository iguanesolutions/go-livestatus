package livestatus

import (
	"errors"
)

var (
	// ErrInvalidQuery represents an invalid query error.
	ErrInvalidQuery = errors.New("invalid query")
	// ErrInvalidType represents an invalid type error.
	ErrInvalidType = errors.New("invalid type")
	// ErrUnknownColumn represents an unknown column error.
	ErrUnknownColumn = errors.New("unknown column")
)
