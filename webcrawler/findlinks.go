package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Crawling: %v\n", err)
		os.Exit(1)
	}

	printText(doc)
	/*
		for _, link := range visit(nil, doc) {
			fmt.Println(link)
		}
	*/
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	/*
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			links = visit(links, c)
		}
	*/
	return links
}

func printText(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
}
