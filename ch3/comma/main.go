// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma2(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
// exercise 3.10
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i , _ := range s {
		buf.WriteString(string(s[i]))
		if (i + 1) % 3 == len(s) % 3 {
			buf.WriteString(",")
		}
	}
	s = buf.String()
	return s[:len(s)-1]
}

// exercise 3.11
func comma2(str string) string {
	if str[0] == '+' {
		str = str[1:]
	}
	i := strings.Index(str, ".")
	s := str[:i] //整数部分
	if str[0] == '-' {
		s = s[1:i]
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	// 获取正负号
	var symbol byte
	if str[0] == '+' || str[0] == '-' {
		symbol = str[0]
		buf.WriteByte(symbol)
		str = str[1:]
	}
	arr := strings.Split(str, ".")

	for i , _ := range arr[0] {
		buf.WriteString(string(arr[0][i]))
		if (i + 1) % 3 == len(arr[0]) % 3 && i != len(arr[0])-1{
			buf.WriteString(",")
		}
	}
	if len(arr) > 1 {
		buf.WriteString(".")
		buf.WriteString(arr[1])
	}

	//s = buf.String()
	//s = s[:len(s)-1]
	//if str[0] == '-' {
	//	s = string('-') + s + str[i:]
	//} else {
	//	s = s + str[i:]
	//}
	return buf.String()
}

//!-
