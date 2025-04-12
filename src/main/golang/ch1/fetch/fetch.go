package fetch

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"net/url"
)

func Fetch(urls []*url.URL) (resp []*http.Response, err []error) {
	resp = make([]*http.Response, len(urls))
	err = make([]error, len(urls))
	for i, urlStruct := range urls {
		resp[i], err[i] = http.Get(urlStruct.String())
	}
	return
}

func ParseHyperLinks(resp *http.Response) (links []string, err error) {
	var node *html.Node
	node, err = html.Parse(resp.Body)
	if err != nil {
		return
	}
	nodeIter := func(yield func(*html.Node) bool) {
		if !yield(node) {
			return
		}
		for n := range node.Descendants() {
			if !yield(n) {
				return
			}
		}
	}
	for node = range nodeIter {
		if node.DataAtom == atom.A {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
	}
	return
}
