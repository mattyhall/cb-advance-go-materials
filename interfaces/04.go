package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type LoggingReader struct{ r io.Reader }

func (r *LoggingReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		fmt.Printf("Got an error: %s\n", err)
		return n, err
	}

	fmt.Printf("Read %d bytes\n", n)

	return n, err
}

func use(r io.Reader) error {
	buf := make([]byte, 1024)

	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	fmt.Printf("Got %v\n", buf[:n])
	return nil
}

func main() {
	f, _ := os.Open("interfaces/04.go")

	buf := []byte{42, 147}
	b := bytes.NewReader(buf)

	use(&LoggingReader{r: f})
	use(&LoggingReader{r: b})
}
