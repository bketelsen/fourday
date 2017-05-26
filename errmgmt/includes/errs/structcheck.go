package errs

import (
	"log"
	"os"
)

// START OMIT
func Login(u, p string) (*Customer, error) {
	cust, err := user.Login(u, p)
	if err != nil {
		return nil, NewError(LoginFailed)
	}
	return cust, nil
}

func main() {
	cust, err := Login(os.Args[1], os.Args[2])
	switch err := err.(type) {
	case nil:
		// OK to proceed
		ProcessTasks()
	case *ApplicationError:
		log.Println("%v:", err.Time, err.Error())
	default:
		// stop execution - unknown error
		log.Fatal(err)
	}
}

// END OMIT
