package main

import "fmt"

func send(ch chan<- int) {
	ch <- 42
}

func recv(ch <-chan int) int {
	return <-ch
}

func main() {
	ch := make(chan int, 1)
	send(ch)
	fmt.Println(recv(ch))
}
