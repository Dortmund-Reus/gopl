// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)
const wordSize = 32 << (^uint(0) >> 63)
//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/wordSize, uint(x%wordSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith 求交集
func (s *IntSet) IntersectWith(t *IntSet) {
	// 首先判断谁比较长
	if len(s.words) > len(t.words) {
		// 如果s比较长，则要注意把后面的全部清零
		s.words = s.words[:len(t.words)]
	}
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			//s.words = append(s.words, tword)
		}
	}
}

//func DifferenceWith 求差集
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			//s.words = append(s.words, tword)
		}
	}
}

//func SymmetricDifference 求并差集,也就是异或
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//func (*IntSet) Len() int      // return the number of elements
//func (*IntSet) Remove(x int)  // remove x from the set
//func (*IntSet) Clear()        // remove all elements from the set
//func (*IntSet) Copy() *IntSet // return a copy of the set

// 返回集合中的元素个数
func (s *IntSet) Len() int {
	cnt := 0
	for _, word := range s.words {
		// 计算word中1的个数
		n := word
		for n > 0 {
			cnt++
			n = n & (n-1)
		}
	}
	return cnt
}

func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		//fmt.Println("找到了！")
		word, bit := x/wordSize, uint(x%wordSize)
		s.words[word] &= 0 << bit
	}
}
//
func (s *IntSet) Clear() {
	for i, _ := range s.words {
		//fmt.Println(word)
		s.words[i] = 0
	}
}
//
func (s *IntSet) Copy() *IntSet {
	//var ret *IntSet = new IntSet{
	//	make([]uint64, len(s.words))
	//}
	ret := new(IntSet)
	ret.words = make([]uint, len(s.words))
	//ret.words = make([]uint64, len(s.words))
	copy(ret.words, s.words)
	return ret
}

func (s *IntSet) AddAll(elements ...int) {
	for _, val := range elements {
		s.Add(val)
	}
}

func (s *IntSet) Elems() []int {
	var res []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, wordSize*i+j)
			}
		}
	}
	return res
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", wordSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string
