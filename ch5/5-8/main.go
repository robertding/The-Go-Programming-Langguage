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
	n := findElementById(doc, "root")
	fmt.Println(n)
}

func startElement(n *html.Node, id string) (isStop bool) {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
	}
	return false
}

func endElement(_ *html.Node, _ string) bool {
	return false
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, id string) bool, id string) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		n := forEachNode(c, pre, post, id)
		if n != nil {
			return n
		}
	}

	if post != nil {
		if post(n, id) {
			return n
		}
	}
	return nil
}

func findElementById(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, startElement, endElement, id)
}
