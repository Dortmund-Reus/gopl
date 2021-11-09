// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse2(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	rotate(s, 2)
	fmt.Println(s)
	//!-slice

	ss := []string{"hello", "hello", "hello", "hey", "hey", "hello", "marco"}
	fmt.Println(ss)
	fmt.Println(distinct(ss))


	str := "我不 喜欢 你的            hat!"
	byteArray := []byte(str)
	fmt.Println(len(byteArray))
	//for _, val := range byteArray {
	//	fmt.Println(val)
	//}
	byteArray = removeSpace(byteArray)
	//fmt.Println(string(removeSpace(byteArray)))
	fmt.Println(string(byteArray))
	reverseByte(byteArray)
	fmt.Println(string(byteArray))

	// Interactive test of reverse.
//	input := bufio.NewScanner(os.Stdin)
//outer:
//	for input.Scan() {
//		var ints []int
//		for _, s := range strings.Fields(input.Text()) {
//			x, err := strconv.ParseInt(s, 10, 64)
//			if err != nil {
//				fmt.Fprintln(os.Stderr, err)
//				continue outer
//			}
//			ints = append(ints, int(x))
//		}
//		reverse(ints)
//		fmt.Printf("%v\n", ints)
//	}
	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// exercise 4.3
func reverse2(p *[6]int) {
	var arr [6]int = *p
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// exercise 4.4
// 为什么这个函数不会改变s?
func rotate(s []int, index int) {
	for i := 0; i < index; i++ {
		t := s[0]
		s = s[1:]
		s = append(s, t)
		fmt.Println(s)
	}
	//return s
}
//!-rev

// exercise 4.5
func distinct(ss []string) []string {
	//tmp := ss
	for i := 1; i < len(ss); i++ {
		if ss[i] == ss[i-1] {
			// 去掉自己
			tmp := ss[:i]
			tmp = append(tmp, ss[i+1:]...)
			//fmt.Println(tmp)
			ss = tmp
			i--
		}
	}
	return ss
}

// exercise 4.6
func removeSpace(str []byte) []byte{
	for i := 1; i < len(str); i++ {
		if unicode.IsSpace(rune(str[i])) && unicode.IsSpace(rune(str[i-1])) {
			// 去掉自己
			tmp := str[:i]
			tmp = append(tmp, str[i+1:]...)
			//fmt.Println(tmp)
			str = tmp
			i--
		}
	}

	return str
}


// exercise 4.7
func reverseByte(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}