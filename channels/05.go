package main

import (
	"fmt"
	"net/http"
	"sync"
)

type DownResult struct {
	url  string
	down bool
}

func checkIfDown(url string) bool {
	res, err := http.Get(url)
	if err != nil {
		return true
	}

	return res.StatusCode != http.StatusOK
}

func worker(input <-chan string, output chan<- DownResult) {
	for {
		url, ok := <-input
		if !ok {
			return
		}

		output <- DownResult{url: url, down: checkIfDown(url)}
	}
}

func main() {
	urls := []string{
		"https://www.couchbase.com", "https://www.bbc.co.uk", "https://www.mattjhall.xyz/blah", "https://www.google.com",
		"https://yahoo.com",
	}

	input := make(chan string, 1)
	output := make(chan DownResult, 1)

	wg := sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			worker(input, output)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < len(urls); i++ {
			out := <-output
			s := "up"
			if out.down {
				s = "down"
			}

			fmt.Printf("%q is %s\n", out.url, s)
		}
	}()

	for _, url := range urls {
		input <- url
	}

	close(input)

	wg.Wait()
}
