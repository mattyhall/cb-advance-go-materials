package main

import (
	"sync"
	"context"
	"fmt"
)

func doWork(ctx context.Context) {
	fmt.Printf("Hello %s\n", ctx.Value("name"))
}

func main() {
	ctx := context.WithValue(context.Background(), "name", "Matt")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		doWork(ctx)
	}()

	wg.Wait()
}
