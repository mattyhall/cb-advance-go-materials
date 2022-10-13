package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		fmt.Printf("Appending %d to channel\n", i)
		ch <- i
		fmt.Printf("len=%d, cap=%d\n", len(ch), cap(ch))
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("Received %d from channel\n", <- ch)
		fmt.Printf("len=%d, cap=%d\n", len(ch), cap(ch))
	}
}
