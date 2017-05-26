package main

import (
	"context"
	"testing"
	"time"
)

var results []string

func BenchmarkSearch(b *testing.B) {

	path := "/Users/bketelsen/src/github.com/golang"
	pattern := "fmt"
	ctx, _ := context.WithTimeout(context.Background(), 10000*time.Millisecond)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		m, err := search(ctx, path, pattern)
		results = m
		if err != nil {
			b.Fatal(err)
		}
	}

}
