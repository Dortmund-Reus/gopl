package main

import (
	"fmt"
	"sort"
)

type customSort struct {
	s []string
}

func (c *customSort) Len() int {return len(c.s)}
func (c *customSort) Less(i, j int) bool {return c.s[i] < c.s[j]}
func (c *customSort) Swap(i, j int) {c.s[i], c.s[j] = c.s[j], c.s[i]}

func NewCustomSort(s []string) *customSort {
	return &customSort{s}
}

func IsPalindrome(s sort.Interface) bool {
	i, j := 0, s.Len()-1
	for ; i < j; {
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	strs := []string{"h", "q", "s", "q", "h"}
	nsc := NewCustomSort(strs)
	if IsPalindrome(nsc) {
		fmt.Println("是回文序列")
	}
}
