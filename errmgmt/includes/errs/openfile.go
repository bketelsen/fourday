package errs

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// START OMIT
func getConfig(path string) ([]byte, error) {
	bb, err := ioutil.ReadFile(path)
	return bb, errors.Wrap(err, "get config")
}

func initialize(path string) error {
	config, err := getConfig(path)
	if err != nil {
		return errors.Wrap(err, "initialize")
	} // process the config file, open database, etc
}

func main() {
	if err := initialize(os.Args[1]); err != nil {
		errors.Print(err)
		switch err := errors.Cause(err).(type) {
		case os.PathError:
			log.Error("Unable to open log file, check location:", err)
		default:
			log.Error("Unknown error:", err)
		}
		os.Exit(-1)
	}
}

// END OMIT
