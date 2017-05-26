package errs

import (
	"fmt"
	"time"
)

// START OMIT
type ErrCode int

var (
	LoginFailed ErrCode = 1000
)

var appErrors = map[ErrCode]string{
	LoginFailed: "Login Failed",
}

type ApplicationError struct {
	Time time.Time
	Code ErrCode
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, appErrors[e.Code])
}
func NewError(code ErrCode) *ApplicationError {
	return &ApplicationError{
		Time: time.Now(),
		Code: code,
	}
}

//END OMIT
