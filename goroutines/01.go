package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting")

	timeout := 5 * time.Second

	go func() {
		time.Sleep(timeout)

		fmt.Printf("We waited %s\n", timeout)
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("We waited three seconds")
}
