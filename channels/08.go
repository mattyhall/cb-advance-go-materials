package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := make(chan int, 1)
	b := make(chan int, 1)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		select {
			case v := <-a:
				fmt.Printf("Got %d from a\n", v)
			case v := <-b:
				fmt.Printf("Got %d from b\n", v)
		}
	}()

	rand.Seed(time.Now().Unix())

	if rand.Float32() >= 0.5 {
		a <- 42
	} else {
		b <- 147
	}

	wg.Wait()
}
