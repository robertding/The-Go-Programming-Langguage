package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

var nodeCount = make(map[string]int)

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
	visit(doc)
	for node, count := range nodeCount {
		fmt.Println(node, count)
	}
}

func visit(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		nodeCount[n.Data]++
	}
	visit(n.NextSibling)
	visit(n.FirstChild)
}
