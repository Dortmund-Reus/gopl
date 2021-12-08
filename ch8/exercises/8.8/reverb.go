// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
"bufio"
"fmt"
	"io"
	"log"
"net"
"runtime"
"strings"
"time"
)


func echo(c net.Conn, shout string, delay time.Duration) {
	buf := make([]byte, 1000)
	runtime.Stack(buf, false)
	fmt.Println(string(buf))
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	//input := bufio.NewScanner(c)
	timeout := 2 * time.Second
	ticker := time.NewTicker(timeout)
	defer func() {
		ticker.Stop()
		c.Close()
	}()
	lines := make(chan string)
	go getInput(c, lines)
	for {
		//fmt.Println(countdown)
		select {
		case text := <-lines:
			ticker.Reset(timeout)
			go echo(c, text, 1*time.Second)
		case <-ticker.C:
			return
		}
	}
	//launch()
}

func getInput(r io.Reader, lines chan string) {

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines <- s.Text()
	}
	// scan will most likely try to read from the connection after it's closed
	// by handleConn. I don't know how to avoid this. Go seems to shun async io
	// in favour of goroutines, so it probably isn't worth avoiding.
	if s.Err() != nil {
		log.Print("scan: ", s.Err())
	}
	//go handleConn(c, input_)
}

//!-

func main() {
	//input_c := make(chan string)
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
		//go getInput(conn)
	}
}

