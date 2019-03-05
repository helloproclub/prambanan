package errors

import (
	"fmt"
)

// Error struct
type Error struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewError constructor like of Error
func NewError(status int, code int, message string) error {
	return &Error{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

// Error implement an interface error in builtin/builtin.go
func (e *Error) Error() string {
	return fmt.Sprintf("[%v] %v", e.Code, e.Message)
}
