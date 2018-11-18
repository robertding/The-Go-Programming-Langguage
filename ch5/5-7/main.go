package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	url := "https://m.zhihu.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get url fail: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s/>\n", depth*2, "", n.Data)
		} else {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

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
