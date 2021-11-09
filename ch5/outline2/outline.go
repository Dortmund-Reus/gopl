// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
	//fmt.Println(expand("foo", doubleString))
}

func outline(url string) error {
	depth := 0
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	// exercise 5.12
	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	endElement := func (n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	forEachNode(doc, startElement, endElement)
	//!-call

	//res := ElementByID(doc, "toc")
	//fmt.Println(res)

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
//var depth int

//func startElement(n *html.Node) {
//	if n.Type == html.ElementNode {
//		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
//		depth++
//	}
//}
//
//func endElement(n *html.Node) {
//	if n.Type == html.ElementNode {
//		depth--
//		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
//	}
//}

// exercise 5.7
//func startElement(n *html.Node) {
//	if n.Type == html.ElementNode {
//		//fmt.Println(len(n.Attr))
//		//fmt.Println(reflect.TypeOf(n.Attr))
//		//生成属性字符串
//		str := ""
//		for _, attr := range n.Attr {
//			str += " "
//			str += attr.Key
//			str += "='"
//			str += attr.Val
//			str += "'"
//		}
//		if len(n.Attr) == 0 {
//			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
//		} else {
//			fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, str)
//		}
//
//		depth++
//	} else if n.Type == html.TextNode {
//		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
//		//depth++
//	} else if n.Type == html.CommentNode {
//		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
//	}
//}
//
//func endElement(n *html.Node) {
//	if n.Type == html.ElementNode {
//		depth--
//		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
//	}
//}

// exercise 5.8
//func forEachNode(n *html.Node, id string, pre func(n *html.Node, id string) bool, post func(n *html.Node)) *html.Node {
//	if pre != nil {
//		if pre(n, id) == true {
//			//fmt.Println("找到了！", n)
//			return n
//		}
//	}
//
//	for c := n.FirstChild; c != nil; c = c.NextSibling {
//		t := forEachNode(c, id, pre, post)
//		if t != nil {
//			return t
//		}
//	}
//
//	if post != nil {
//		post(n)
//	}
//	return nil
//	//return nil
//}

//func ifEqual(n *html.Node, id string) bool {
//	if n.Type == html.ElementNode {
//		for _, attr := range n.Attr {
//			if attr.Key == "id" && attr.Val == id {
//				return true
//			}
//		}
//		//depth--
//		//fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
//	}
//	return false
//}
//
//func ElementByID(doc *html.Node, id string) *html.Node {
//
//	return forEachNode(doc, id, ifEqual, endElement)
//}

func doubleString(s string) string {
	s += s
	return s
}

// exercise 5.9
func expand(s string, f func(string) string) string {
	return f(s)
}

//!-startend
