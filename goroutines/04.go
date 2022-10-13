package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	unsorted := []int{4, 1, 0, 5, 3}

	wg := sync.WaitGroup{}
	for _, n := range unsorted {
		i := n

		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Printf("%d ", i)
		}()
	}

	wg.Wait()
	fmt.Println("")
}
