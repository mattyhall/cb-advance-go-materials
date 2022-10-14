package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
)

type DownResult struct {
	url  string
	down bool
}

func checkIfDown(ctx context.Context, url string) (bool, error) {
	cli := http.DefaultClient

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false, err
	}

	res, err := cli.Do(req)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			return true, nil
		}

		return false, err
	}

	return res.StatusCode != http.StatusOK, nil
}

func worker(ctx context.Context, input <-chan string, output chan<- DownResult) error {
	for {
		var (
			url string
			ok  bool
		)

		select {
		case url, ok = <-input:
			if !ok {
				return nil
			}

		case <-ctx.Done():
			return ctx.Err()
		}

		down, err := checkIfDown(ctx, url)
		if err != nil {
			return err
		}

		select {
		case output <- DownResult{url: url, down: down}:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	urls := []string{
		"https://www.couchbase.com", "https://www.bbc.co.uk", "https://www.mattjhall.xyz/blah", "https://www.google.com",
		"https://yahoo.com",
	}

	input := make(chan string, 1)
	output := make(chan DownResult, 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := worker(ctx, input, output); err != nil {
				fmt.Printf("Got error from worker: %s\n", err)
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < len(urls); i++ {
			select {
			case out := <-output:
				s := "up"
				if out.down {
					s = "down"
				}

				fmt.Printf("%q is %s\n", out.url, s)
			case <-ctx.Done():
				return
			}
		}
	}()

	for i, url := range urls {
		if i == 4 {
			cancel()
		}

		select {
		case input <- url:
		case <-ctx.Done():
			break
		}
	}

	close(input)

	wg.Wait()
}
