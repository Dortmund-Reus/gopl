package main

import (
	"fmt"
	"sync"
)

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// pc[i] is the population count of i.
var pc [256]byte

var loadPCOnce sync.Once

func Myinit() {
	fmt.Println("初始化")
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	loadPCOnce.Do(Myinit)
	sum := 0
	for x != 0 {
		sum ++
		x = x & (x-1)
	}
	return sum
}

func main() {
	fmt.Println(PopCount(991))
	fmt.Println(PopCount(919))
	fmt.Println(PopCount(199))
	fmt.Println(PopCount(1199))
	fmt.Println(PopCount(1919))
	fmt.Println(PopCount(1991))

}

//!-
