package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1

	n := <-ch

	fmt.Println(n)
}
