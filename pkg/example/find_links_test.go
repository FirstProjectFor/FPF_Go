package demo

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestFindLinks(t *testing.T) {

	for _, url := range os.Args[1:] {
		fmt.Printf("get link from %s", url)
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks: %v \n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}

}

func findLinks(url string) ([]string, error) {
	res, err := http.Get(url)
	if nil != err {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("getting %s： %s", url, res.Status)
	}

	doc, err := html.Parse(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s an HTML: %v", url, err)
	}

	return visit(nil, doc), nil
}

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
