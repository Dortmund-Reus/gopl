// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": 		{"data structures"},
	"calculus":   		{"linear algebra"},
	"linear algebra":   {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	//for i, course := range topoSort(prereqs) {
	//	fmt.Printf("%d:\t%s\n", i+1, course)
	//}

	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
//func topoSort(m map[string][]string) []string {
//	var order []string
//	seen := make(map[string]bool)
//	var visitAll func(items []string)
//
//	visitAll = func(items []string) {
//		for _, item := range items {
//			//fmt.Println(item)
//			if !seen[item] {
//				seen[item] = true
//				visitAll(m[item])
//				order = append(order, item) //递归调用，所以顺序其实是对的
//				//fmt.Println(order)
//			}
//		}
//	}
//
//	var keys []string
//	for key := range m {
//		keys = append(keys, key)
//	}
//
//	sort.Strings(keys)
//	//fmt.Println(keys)
//	visitAll(keys)
//	return order
//}

// exercise 5.10
//func topoSort(m map[string][]string) []string {
//	var order []string
//	seen := make(map[string]bool)
//	var visitAll func(items []string)
//
//	visitAll = func(items []string) {
//		for _, item := range items {
//			if !seen[item] {
//				seen[item] = true
//				visitAll(m[item])
//				order = append(order, item) //递归调用，所以顺序其实是对的
//			}
//		}
//	}
//	for key, _ := range m {
//		visitAll([]string{key})
//	}
//	return order
//}

// exercise 5.11

// 返回s在slice中的下标，如果没有的话，返回0和错误信息
func index(s string, slice []string) (int, error) {
	for i, v := range slice {
		if s == v {
			return i, nil
		}
	}
	return 0, fmt.Errorf("not found")
}

func topoSort(m map[string][]string) (order []string, err error) {
	resolved := make(map[string]bool)
	var visitAll func([]string, []string)

	visitAll = func(items []string, parents []string) {
		fmt.Println(parents)
		for _, v := range items {
			vResolved, seen := resolved[v] //读取map是会返回两个值的！
			fmt.Println(vResolved, seen)
			//test1, test2 := 1
			//fmt.Println(test1, test2)
			if seen && !vResolved {
				// seen为true而vResolved为false，说明出问题了
				// 这里的parents容易引起误解，实际上parents中的应当是高级课程，而当前的才是初级课程
				start, _ := index(v, parents) // Ignore error since v has to be in parents.
				err = fmt.Errorf("cycle: %s", strings.Join(append(parents[start:], v), " -> "))
			}
			if !seen {
				resolved[v] = false
				visitAll(m[v], append(parents, v))
				resolved[v] = true
				order = append(order, v)
			}
		}
	}

	for k := range m {
		if err != nil {
			return nil, err
		}
		visitAll([]string{k}, nil)
	}
	return order, nil
}

//!-main
