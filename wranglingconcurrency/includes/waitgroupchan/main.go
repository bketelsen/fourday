package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

// START OMIT
func main() {
	uppers := make(chan string, len(os.Args)-1)
	var wg sync.WaitGroup
	for i := 1; i < len(os.Args); i++ { // arg[0] is exe name
		i := i // WHY?  closure and goroutine!
		wg.Add(1)
		go func() {
			defer wg.Done()
			uppers <- strings.ToUpper(os.Args[i])
		}()
	}
	wg.Wait()
	close(uppers)
	for u := range uppers {
		fmt.Println(u)
	}
}

// END OMIT
