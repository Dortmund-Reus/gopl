package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

var done = make(chan struct{})

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("http://gopl.io") }()
	go func() { responses <- request("http://www.github.com") }()
	go func() { responses <- request("http://gopl.io") }()
	return <-responses // return the quickest response
}

func fetch(url string) (content string, n int64, err error) {

	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	//local := path.Base(resp.Request.URL.Path)
	//if local == "/" {
	//	local = "index.html"
	//}
	//f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	//var s string
	buff := new(bytes.Buffer)
	n, err = io.Copy(buff, resp.Body)
	//s := buff.String()
	// Close file, but prefer error from Copy, if any.
	//if closeErr := f.Close(); err == nil {
	//	err = closeErr
	//}
	if cancelled() {
		return
	}
	close(done)
	return buff.String(), n, err
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func request(hostname string) (content string) {

	content, _, err := fetch(hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s: %v\n", hostname, err)
	}
	//fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	//fmt.Println(content)

	fmt.Println("关闭通道的url是:", hostname)
	return content
}

func main() {
	content := mirroredQuery()
	fmt.Println(content)
}
