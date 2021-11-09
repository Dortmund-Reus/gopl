package main

import (
	"fmt"
)

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

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
//
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



func deps(course string) []string {
	fmt.Println(course)
	return prereqs[course]
}

func main() {
	var course string
	for course = range prereqs { // get random key
		break
	}
	//for course := range prereqs {
	//	breadthFirst(deps, []string{course})
	//}
	breadthFirst(deps, []string{course})
}