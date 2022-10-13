package main

import "fmt"

func main() {
	a := make(chan int, 1)
	b := make(chan int, 1)

	go func() {
    <-a
    b <- 1
	}()

	fmt.Println(<-b)
	a <- 2
}
