package errs

import (
	"fmt"
	"os"
)

// START OMIT

func Login(u, p string) error {
	available := db.Alive() //bool
	if !available {
		return ErrDatabaseUnavailable
	}
	err := user.Login(u, p)
	if err != nil {
		return ErrInvalidLogin
	}
	return nil
}

func main() {
	err := Login(os.Args[1], os.Args[2])
	if err != nil {
		if err == ErrDatabaseUnavailable {
			fmt.Println("Database Not Available, try later")
		}
		if err == ErrInvalidLogin {
			fmt.Println("Bad Login")
		}
		os.Exit(-1)
	}
}

//END OMIT
