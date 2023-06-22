package webp

import "errors"

var (
	ErrWrongFile         = errors.New("not a webp file")
	ErrMissingWebPHeader = errors.New("webp: missing WebP chunk header")
	ErrMissingHeader     = errors.New("missing header")
)
