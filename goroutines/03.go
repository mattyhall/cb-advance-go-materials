package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		j := i

		wg.Add(1)
		go func() {
			defer wg.Done()

			time.Sleep(time.Duration(j) * time.Second)
			fmt.Printf("We slept for %d secs\n", j)
		}()
	}

	wg.Wait()
}
