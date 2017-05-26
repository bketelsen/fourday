package errs

import (
	"io/ioutil"
	"log"
)

// START OMIT
func GetConfig() error {
	bb, err := ioutil.ReadFile("/etc/app/config")
	if err != nil {
		log.Println("Couldn't read config:", err)
		return nil, err
	}
	return bb, nil
}

// END OMIT
