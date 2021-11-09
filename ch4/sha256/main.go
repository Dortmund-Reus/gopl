// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

var pc [256]byte

// exercise 4.2
func main() {
	//var c1, c2 [32]uint8
	if(len(os.Args) == 1) {
		c1 := sha256.Sum256([]byte("x"))
		c2 := sha256.Sum256([]byte("X"))
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	} else if os.Args[1] == "sha384" {
		c1 := sha512.Sum384([]byte("x"))
		c2 := sha512.Sum384([]byte("X"))
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	} else if os.Args[1] == "sha512" {
		c1 := sha512.Sum512([]byte("x"))
		c2 := sha512.Sum512([]byte("X"))
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	}

	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
	//sum := 0
	//for i , _ := range c1 {
	//	sum += PopCount(c1[i] ^ c2[i])
	//}
	//fmt.Println(sum)
}
// exercise 4.1
func PopCount(x uint8) int {
	sum := 0
	for x != 0 {
		sum ++
		x = x & (x-1)
	}
	return sum
}

// exercise 4.2

//!-
