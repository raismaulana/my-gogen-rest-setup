package apperror

import (
	"fmt"
	"strconv"
	"strings"
)

// Error is defined as this sample
// TransitionError ErrorType = "ER1043 Transition from %s to %s is not allowed"
//
// All Error is registered in application/apperror/error_enum.go file
//
// TransitionError
// 		is the Error Enum
// ER1000
// 		is the Error Code. You may customize your own code format here
// Transition from %s to %s is not allowed
// 		is the message with optional formatted variable

// ErrorType must not modified
type ErrorType string

type ErrorWithCode interface {
	error
	Code() string
	CodeInt() int
}

const errorCodePrefix = "ER"

// Error return the only message
func (u ErrorType) Error() string {
	s := string(u)
	if strings.HasPrefix(s, errorCodePrefix) {
		i := strings.Index(s, " ")
		return s[i+1:]
	}
	return s
}

// Code return the only code
func (u ErrorType) Code() string {
	s := string(u)
	if strings.HasPrefix(s, errorCodePrefix) {
		i := strings.Index(s, " ")
		return s[:i]
	}
	return ""
}

// Code return the only code int type
func (u ErrorType) CodeInt() int {
	s := string(u)
	if strings.HasPrefix(s, errorCodePrefix) {
		i := strings.Index(s, " ")
		code := strings.TrimPrefix(s[:i], errorCodePrefix)
		codeInt, err := strconv.Atoi(code)
		if err != nil {
			return 500
		}
		return codeInt
	}
	return 500
}

// Var add generic variable value to the error message
// for example you have
// UserNotFoundError ErrorType = "ER1092 User with name %s is not found"
// Then you can insert the name
// UserNotFoundError.Var("mirza") --> "User with name mirza is not found"
func (u ErrorType) Var(params ...interface{}) ErrorType {
	return ErrorType(fmt.Sprintf(u.String(), params...))
}

// String return the error as it is
func (u ErrorType) String() string {
	return string(u)
}

// Extract error code int apperror
func GetErrorCode(err error) int {
	et, ok := err.(ErrorWithCode)
	if !ok {
		return 500
	}
	return et.CodeInt()
}
