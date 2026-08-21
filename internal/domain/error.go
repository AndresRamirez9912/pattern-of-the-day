package domain

// Type assertion Error must implement error interface
var _ error = &Error{}

// Error represents a domain error
type Error struct {
	Message string
	Code    int
}

// NewError creates a new error with the given message and code
func NewError(message string, code int) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

// Error returns the error message for SimpleError
func (e Error) Error() string {
	return e.Message
}

// Is implement the comparable interface for the Error type
func (e Error) Is(other error) bool {
	// Check if the other error is nil
	if other == nil {
		return false
	}

	// Check if the other error is of type Error
	// It must have equal code to be considered equal
	otherErr, ok := other.(*Error)
	if ok {
		return e.Code == otherErr.Code
	}

	return false
}
