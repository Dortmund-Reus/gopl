package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var n sync.WaitGroup
	//n.Add(1)
	ch1 := make(chan string)
	ch2 := make(chan string)

	var i int64
	start := time.Now()
	//n.Add(1)
	go func() {

		//defer n.Done()
		for {
			i++
			select {
			case ch1<- "ping-pong":
				//do nothing
			case <-ch2:
				//do nothing
			}
		}
		//n.Done()
	}()
	n.Add(1)
	go func() {

		for {
			i++
			select {
			case ch2<- "ping-pong":
				//do nothing
			case <-ch1:
				//do nothing
			}
		}
	}()

	//go func() {
	//
	//}()
	//for {
	//
	//}

	//n.Wait()
	//close(ch1)
	//close(ch2)
	//fmt.Println("hahahah")

	//q := make(chan int)
	//
	//go func() {
	//	q <- 1
	//	for {
	//		i++
	//		q <- <-q
	//	}
	//}()
	//go func() {
	//	for {
	//		q <- <-q
	//	}
	//}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println(float64(i)/float64(time.Since(start))*1e9, "round trips per second")
}
