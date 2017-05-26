package errs

import "errors"

// START OMIT
var (
	// ErrInvalidLogin is returned when a bad username or password
	// is used to attempt a login to the application
	ErrInvalidLogin = errors.New("Invalid Login")
	// ErrDBUnavailable is returned with the database can't be contacted
	ErrDBUnavailable = errors.New("Database Unavailable")
)

// END OMIT
