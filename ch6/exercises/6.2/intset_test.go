// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "fmt"

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	new_x := x.Copy()
	fmt.Println(new_x)// "{1 9 42 144}"
	fmt.Println(x.Len()) // 4
	x.Remove(144)
	fmt.Println(&x)

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	x.Clear()
	fmt.Println(&x)
	fmt.Println(new_x)
	new_x.AddAll(2,3,5)
	fmt.Println(new_x)
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// {1 9 42 144}
	// 4
	// {1 9 42}
	// true false
	// {}
	// {1 9 42 144}
	// {1 2 3 5 9 42 144}
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}
