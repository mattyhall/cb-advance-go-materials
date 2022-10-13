package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting")

	timeout := 5 * time.Second

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		time.Sleep(timeout)

		fmt.Printf("We waited %s\n", timeout)
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("We waited 3s")

	wg.Wait()
}
