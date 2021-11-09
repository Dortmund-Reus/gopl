// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import (
	"fmt"
	"math"
)

//!+
func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	} else {
		res := math.MaxInt32
		for _, val := range vals {
			if res > val {
				res = val
			}
		}
		return res
	}
}

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	} else {
		res := math.MinInt32
		for _, val := range vals {
			if res < val {
				res = val
			}
		}
		return res
	}
}

func main() {
	test_nums := []int{1, 2, 4, 5}
	fmt.Println(min(test_nums...))
	fmt.Println(max(test_nums...))
	fmt.Println(min(1))
	fmt.Println(max(1))
}
