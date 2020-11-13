package gotils

import (
	"errors"
	"fmt"
)

// ErrNotFound generic sentinel not found error
var ErrNotFound = errors.New("not found")

// UserError let's you set a separate error message intended for the end user.
// See UserErrorf for creating one.
// eg:
//   var e core.UserError
//   if errors.As(err, &e) {
// 	   sendToUser(e.UserError())
//     if (e.Unwrap() != nil){
//       log(e)
//     }
// 	 } else {
//     sendGenericInternalErrorToUser()
// 	   log(e)
//   }
type UserError interface {
	error
	UserError() string
}

type userError struct {
	userMsg string
	root    error
}

func (ue *userError) Error() string {
	if ue.root != nil {
		return ue.root.Error()
	}
	return ue.UserError()
}

func (ue *userError) UserError() string {
	return ue.userMsg
}
func (ue *userError) Unwrap() error {
	return ue.root
}

// func uf(format string, a ...interface{}) UserError {
// 	return &userError{userMsg: fmt.Sprintf(format, a...)}
// }

// UserErrorf returns a new UserError
// rootErr can be nil which means this will ONLY be a UserError, not an internal error also,
// which can be useful for returning validation messages and things.
func UserErrorf(rootErr error, format string, a ...interface{}) error {
	return &userError{root: rootErr, userMsg: fmt.Sprintf(format, a...)}
}

// DetailedError let's you add a more detailed message to an error
type DetailedError struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func (e *DetailedError) Error() string {
	return e.Message
}

// HTTPError allows you to add a status code
type HTTPError interface {
	error
	Code() int
}

type httpError struct {
	msg  string
	code int
}

// NewHTTPError create an HTTPError
func NewHTTPError(msg string, code int) HTTPError {
	return &httpError{msg: msg, code: code}
}

func (e *httpError) Error() string {
	return e.msg
}
func (e *httpError) Code() int {
	return e.code
}
