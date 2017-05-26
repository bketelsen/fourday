package throughput

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

type Config struct {
	Listen string
}

func GetConfig(r io.ReadCloser) (*Config, error) {
	var config Config
	defer r.Close()
	dec := json.NewDecoder(r)
	err := dec.Decode(&config)
	return &config, errors.Wrap(err, "decoding config")
}

func (c *Config) Save(w io.WriteCloser) error {
	enc := json.NewEncoder(w)
	defer w.Close()
	return errors.Wrap(enc.Encode(c), "encoding config")
}
