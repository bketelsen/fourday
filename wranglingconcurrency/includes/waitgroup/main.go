package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i < len(os.Args); i++ { // arg[0] is exe name
		i := i // WHY?  closure and goroutine!
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(strings.ToUpper(os.Args[i]))
		}()
	}
	wg.Wait()
}
