// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 125.

// Findlinks2 does an HTTP GET on each URL, parses the
// result as HTML, and prints the links within it.
//
// Usage:
//	findlinks url ...
package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

// visit appends to links each link found in n, and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

//!+
func main() {

	//str := "Hello hello  hello  hhell"
	//fmt.Println(countWords(str))
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
		words, images, _ := CountWordsAndImages(url)
		fmt.Printf("words = %d, images = %d\n", words, images)
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}


// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

// exercise 5.5
func countWordsAndImages(n *html.Node) (words, images int) {

	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		//words++
		//fmt.Println(reflect.TypeOf(n.Data))
		words += countWords(n.Data)
		//fmt.Println(n.Data)
	}
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
			//fmt.Println(n)
		}
	}
	word, image := countWordsAndImages(n.FirstChild)
	words += word
	images += image
	word, image = countWordsAndImages(n.NextSibling)
	words += word
	images += image
	return
}

func countWords(text string) int {

	n := 0
	scan := bufio.NewScanner(strings.NewReader(text))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return n

}

//!-
