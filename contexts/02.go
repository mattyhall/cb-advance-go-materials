package main

import (
	"time"
	"sync"
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		select {
			case <-time.After(5 * time.Second):
				fmt.Println("Slept for five seconds")
			case <-ctx.Done():
				fmt.Printf("Timed out (%s)\n", ctx.Err())
		}
	}()

	wg.Wait()
}
