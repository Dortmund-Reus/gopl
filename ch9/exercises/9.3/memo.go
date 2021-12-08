// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 276.

// Package memo provides a concurrency-safe memoization a function of
// a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a Mutex.
package main

import (
	"fmt"
	//"gopl.io/ch9/memotest"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	//"testing"
	"time"
)

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//!+
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

var done = make(chan struct{})
//var httpGetBody = memotest.HTTPGetBody
var HTTPGetBody = httpGetBody

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func testConcurrent() {
	//var memo Memo
	m := New(httpGetBody)
	Concurrent(m)
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		if cancelled() {
			fmt.Println("取消操作！")
			return
		}
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}

//!+httpRequestBody
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//!-httpRequestBody



func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://www.baidu.com",
			"https://www.163.com",
			"https://www.sina.com",
			"http://www.sohu.com",
			//"http://www.t114.cn/",
			"http://www.21-sun.com/",
			"http://china.machine365.com/",
			//"www.chinapackexpo.com",
			"http://www.mei.net.cn/",
			"http://www.mechnet.com.cn/ ",
			//"http://www.newsccn.com/",
			"http://www.lmjx.net/",
			"http://www.d1cm.com/",
			"http://www.hc360.com/",
			"https://www.baidu.com",
			"https://www.163.com",
			"https://www.sina.com",
			"http://www.sohu.com",
			//"http://www.t114.cn/",
			"http://www.21-sun.com/",
			"http://china.machine365.com/",
			//"www.chinapackexpo.com",
			"http://www.mei.net.cn/",
			"http://www.mechnet.com.cn/ ",
			//"http://www.newsccn.com/",
			"http://www.lmjx.net/",
			"http://www.d1cm.com/",
			"http://www.hc360.com/",
			"https://www.baidu.com",
			"https://www.163.com",
			"https://www.sina.com",
			"http://www.sohu.com",
			//"http://www.t114.cn/",
			"http://www.21-sun.com/",
			"http://china.machine365.com/",
			//"www.chinapackexpo.com",
			"http://www.mei.net.cn/",
			"http://www.mechnet.com.cn/ ",
			//"http://www.newsccn.com/",
			"http://www.lmjx.net/",
			"http://www.d1cm.com/",
			"http://www.hc360.com/",
			"https://www.baidu.com",
			"https://www.163.com",
			"https://www.sina.com",
			"http://www.sohu.com",
			//"http://www.t114.cn/",
			"http://www.21-sun.com/",
			"http://china.machine365.com/",
			//"www.chinapackexpo.com",
			"http://www.mei.net.cn/",
			"http://www.mechnet.com.cn/ ",
			//"http://www.newsccn.com/",
			"http://www.lmjx.net/",
			"http://www.d1cm.com/",
			"http://www.hc360.com/",
			//"https://golang.org",
			//"https://godoc.org",
			//"https://play.golang.org",
			//"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string) (interface{}, error)
}

func Sequential(m M) {
	//!+seq
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
	//!-seq
}

/*
//!+conc
	m := memo.New(httpGetBody)
//!-conc
*/

func Concurrent(m M) {
	//!+conc
	var n sync.WaitGroup

	for url := range incomingURLs() {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			//for range fileSizes {
			//	// Do nothing.
			//}
			fmt.Println("停止！")
			return
		default:
			n.Add(1)
			go func(url string) {
				defer n.Done()
				start := time.Now()
				value, err := m.Get(url)
				if err != nil {
					log.Print(err)
					return
				}
				fmt.Printf("%s, %s, %d bytes\n",
					url, time.Since(start), len(value.([]byte)))
			}(url)
		}

	}
	n.Wait()
	//!-conc
}

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()
	testConcurrent()
}

//!-
