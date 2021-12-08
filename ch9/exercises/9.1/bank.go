// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package main

import "fmt"

type withdraw_s struct {
	amount int
	res chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan withdraw_s)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	s := withdraw_s{amount, make(chan bool)}
	withdraw <- s
	return <- s.res
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case s := <-withdraw:
			if balance < s.amount {
				s.res <- false
				continue
			} else {
				balance -= s.amount
				s.res <- true
			}
		case balances <- balance:
			// do nothing
		}

	}
}

//func init() {
//	go teller() // start the monitor goroutine
//}

func main() {

	done := make(chan struct{})
	go teller()

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	go func() {
		fmt.Println("Withdraw(50) = ", Withdraw(50))
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	go func() {
		fmt.Println("Withdraw(500) = " , Withdraw(500))
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	<-done
	<-done
}

//!-
