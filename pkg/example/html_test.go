package demo

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"testing"
)

func TestHtml(t *testing.T) {
	f := flag.String("name", "xiaotian", "user name")
	fmt.Println(&f)
}

func visist(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visist(links, c)
	}

	return links
}
