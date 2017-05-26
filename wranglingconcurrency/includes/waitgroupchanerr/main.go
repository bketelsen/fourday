package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

// START OMIT
type result struct {
	number int
	err    error
}

func main() { // convert input strings into integers
	results := make(chan result, len(os.Args)-1)
	var wg sync.WaitGroup
	for i := 1; i < len(os.Args); i++ { // arg[0] is exe name
		i := i // WHY?  closure and goroutine!
		wg.Add(1)
		go func() {
			defer wg.Done()
			num, err := strconv.Atoi(os.Args[i])
			results <- result{number: num, err: err}
		}()
	}
	wg.Wait()
	close(results)
	for r := range results {
		fmt.Println(r.number, r.err)
	}
}

// END OMIT
