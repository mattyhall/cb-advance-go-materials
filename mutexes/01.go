package main

import (
	"sync"
	"fmt"
)

type Counter struct { n int64 }

func (c *Counter) Inc() { c.n++ }

func main() {
	counter := Counter{}

	wg := sync.WaitGroup{}

	fn := func() {
		defer wg.Done()

		for i := 0; i < 10000; i++ {
			counter.Inc()
		}
	}

	wg.Add(2)
	go fn()
	go fn()

	wg.Wait()

	fmt.Println(counter.n)
}
