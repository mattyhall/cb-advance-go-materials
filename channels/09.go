package main

import (
	"sync"
	"fmt"
)

func main() {
	a := make(chan int, 1)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		select {
			case v := <-a:
				fmt.Printf("Got %d from a\n", v)
			default:
				fmt.Println("Could not receive on a")
		}
	}()

	wg.Wait()
}
