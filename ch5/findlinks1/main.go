// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	//m := make(map[string]int)
	////visit(nil, doc)
	//countElements(m, doc)
	//
	//for key, val := range m {
	//	fmt.Println(key, val)
	//}

	//printText(doc)

	visitNew(nil, doc)
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) {
	// exerise 5.1
	if n == nil {
		return
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
				links = append(links, a.Val)
			}
		}
	}

	visit(links, n.FirstChild)
	visit(links, n.NextSibling)

	//fmt.Println(666)
	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	links = visit(links, c)
	//}
}

// exercise 5.2
func countElements(m map[string]int, n *html.Node) {
	// exerise 5.1
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		m[n.Data]++
	}

	countElements(m, n.FirstChild)
	countElements(m, n.NextSibling)
}

//exercise 5.3
func printText(n *html.Node) {
	// exerise 5.1
	if n == nil {
		return
	}

	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	//if n.Type == html.ElementNode {
	//	m[n.Data]++
	//}

	printText(n.FirstChild)
	printText(n.NextSibling)
}

//exercise 5.4
func visitNew(links []string, n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					links = append(links, a.Val)
				}
			}
		} else if n.Data == "script" {
			fmt.Println(n.Attr)
		}

	}

	visitNew(links, n.FirstChild)
	visitNew(links, n.NextSibling)
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
