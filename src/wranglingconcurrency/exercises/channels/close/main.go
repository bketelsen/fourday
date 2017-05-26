package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	a := make(chan string)
	go func() {
		for i := 1; i < len(os.Args); i++ { // arg[0] is exe name
			a <- strings.ToUpper(os.Args[i])
		}
		close(a)
	}()
	for s := range a {
		fmt.Println(s)
	}
}
