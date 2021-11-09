package main

import (
	"bytes"
	"fmt"
	"io"
)

type ByteCounter struct {
	w io.Writer
	written int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.written += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &ByteCounter{w, 0}
	return c.w, &c.written
}

func main() {
	b := &bytes.Buffer{}
	c, n := CountingWriter(b)
	data := []byte("hi there")
	c.Write(data)
	if *n != int64(len(data)) {
		fmt.Errorf("%d != %d", n, len(data))
	}
	fmt.Println("success!")
}