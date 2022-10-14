package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		select {
			case <-time.After(2 * time.Second):
				fmt.Println("Five seconds elapsed")
			case <-ctx.Done():
				fmt.Println("Cancelled")
				return
		}
	}()

	rand.Seed(time.Now().Unix())

	if rand.Float32()>= 0.5 {
		fmt.Println("Immediately cancelling")
		cancel()
	}

	wg.Wait()
}
