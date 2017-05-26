package errs

import "io/ioutil"

// START OMIT
func GetConfig() ([]byte, error) {
	return ioutil.ReadFile("/etc/app/config")
}

// END OMIT
