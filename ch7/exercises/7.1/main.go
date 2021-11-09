// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

//type WordCounter int
//
//func (c *WordCounter) Write(p []byte) (int, error) {
//	words := strings.Split(string(p), " ")
//	return len(words), nil
//}

type LineCounter struct {
	lines int
}

func (c *LineCounter) N() int {
	return c.lines
}

func (c *LineCounter) Write(p []byte) (int, error) {
	//cnt := 0
	for _, val := range p {
		if val == '\n' {
			c.lines++
		}
	}
	return len(p), nil
}

type WordCounter struct {
	words  int
	inWord bool
}

// 统计下一个word到来之前的空格数量
func leadingSpaces(p []byte) int {
	count := 0
	cur := 0
	for cur < len(p) {
		r, size := utf8.DecodeRune(p[cur:])
		if !unicode.IsSpace(r) {
			return count
		}
		cur += size
		count++
	}
	return count
}

// 统计下一个空格到来之前的字符数量
func leadingNonSpaces(p []byte) int {
	count := 0
	cur := 0
	for cur < len(p) {
		r, size := utf8.DecodeRune(p[cur:])
		if unicode.IsSpace(r) {
			return count
		}
		cur += size
		count++
	}
	return count
}

// A !IsSpace() -> IsSpace() transition is counted as a word.
//
// I couldn't figure out how to use bufio.ScanWords without either
// double-counting words split across buffer boundaries, giving incorrect
// intermediate counts, or doing some really awkward buffer manipulation.
func (c *WordCounter) Write(p []byte) (n int, err error) {
	cur := 0
	n = len(p)
	for {
		spaces := leadingSpaces(p[cur:])
		cur += spaces
		if spaces > 0 {
			c.inWord = false
		}
		if cur == len(p) {
			return
		}
		if !c.inWord {
			c.words++
		}
		c.inWord = true
		cur += leadingNonSpaces(p[cur:])
		if cur == len(p) {
			return
		}
	}
}

func (c *WordCounter) N() int {
	return c.words
}

func (c *WordCounter) String() string {
	return fmt.Sprintf("%d", c.words)
}
//!-bytecounter

func main() {
	//!+main
	//var c ByteCounter
	//c.Write([]byte("hello"))
	//fmt.Println(c) // "5", = len("hello")
	//
	//c = 0 // reset the counter
	//var name = "Dolly"
	//fmt.Fprintf(&c, "hello, %s", name)
	//fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main



	var lc LineCounter
	p := []byte("one\ntwo\nthree\n")
	length, err := lc.Write(p)
	if length != len(p) {
		fmt.Errorf("写入错误\n")
	}
	if err != nil {
		fmt.Errorf("不知道哪里出错了\n")
	}
	if lc.N() != 3 {
		fmt.Errorf("统计行数出错\n")
	}
	fmt.Println("success!")


	var wc WordCounter
	p2 := []byte("hello halo haha ha")
	length2, err := wc.Write(p2)
	if length2 != len(p2) {
		fmt.Errorf("写入错误\n")
	}
	if err != nil {
		fmt.Errorf("something wrong")
	}
	if wc.N() != 4 {
		fmt.Errorf("统计单词数量出错\n")
	}
	fmt.Println("success!")
}
