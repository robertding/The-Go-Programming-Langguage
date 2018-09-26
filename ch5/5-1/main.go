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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links = visit(links, n.NextSibling)
	return visit(links, n.FirstChild)
}
