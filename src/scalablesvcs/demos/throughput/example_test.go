package throughput

import (
	"flag"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

var useFile *bool

func init() {
	useFile = flag.Bool("usefile", false, "write to a file")
	flag.Parse()

}

type testWriter struct {
	b []byte
}

func (tw *testWriter) Close() error {
	return nil
}

func (tw *testWriter) Write(p []byte) (int, error) {
	for _, by := range p {
		tw.b = append(tw.b, by)
	}
	return len(p), nil
}

func TestWrite(t *testing.T) {
	c := &Config{
		Listen: ":8080",
	}
	var buf []byte
	tester := &testWriter{b: buf}
	if err := c.Save(tester); err != nil {
		t.Error(err)
	}
	want := `{"Listen":":8080"}` + "\n"
	if got := string(tester.b); got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func BenchmarkWrite(b *testing.B) {

	c := &Config{
		Listen: ":8080",
	}
	var w io.WriteCloser
	if *useFile {
		t := strconv.Itoa(rand.Intn(b.N))
		path := filepath.Join(os.TempDir(), t+".json")
		f, err := os.Create(path)
		if err != nil {
			b.Error(err)
		}
		w = f
	} else {

		var buf []byte
		tester := &testWriter{b: buf}
		w = tester
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Save(w)
	}
}
