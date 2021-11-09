package main

import (
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	// 层序遍历
	// 需要一个栈
	str := ""
	//stack := make([]*tree, 20)
	var stack []*tree
	stack = append(stack, t)
	for len(stack) != 0 {
		tmp := stack[0]
		str += strconv.Itoa(tmp.value)
		str += " "
		if tmp.left != nil {
			stack = append(stack, tmp.left)
		}
		if tmp.right != nil {
			stack = append(stack, tmp.right)
		}
		stack = stack[1:]
	}
	return str
}

func main() {

	var root *tree
	data := make([]int, 50)
	for i := range data {
		data[i] = i
	}
	for _, v := range data {
		root = add(root, v)
	}
	fmt.Println(root)
}