package errs

import "fmt"

// START OMIT

type Customer struct {
	ID     int
	Name   string
	ErrMsg string
}

func (c *Customer) Error() string {
	if c.Error != nil {
		return fmt.Sprintf("Error: %s", c.ErrMsg)
	}
	return nil
}

//END OMIT
