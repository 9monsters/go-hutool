package jwt

import "errors"

var (
	ErrInvalidKey      = errors.New("key is invalid")
	ErrInvalidKeyType  = errors.New("key is of invalid type")
	ErrHashUnavailable = errors.New("the requested hash function is unavailable")
)

const (
	ValidationErrorMalformed uint32 = 1 << iota
	ValidationErrorUnverifiable
	ValidationErrorSignatureInvalid
)

type ValidationError struct {
	Inner  error
	Errors uint32
	text   string
}

// Error is the implementation of the err interface.
func (e ValidationError) Error() string {
	if e.Inner != nil {
		return e.Inner.Error()
	} else if e.text != "" {
		return e.text
	} else {
		return "token is invalid"
	}
}

func NewValidationError(errorText string, errorFlags uint32) *ValidationError {
	return &ValidationError{
		text:   errorText,
		Errors: errorFlags,
	}
}
