package main

import (
	"examples/ch1/fetch"
	"fmt"
	"net/url"
)

func main() {
	urlStrings := []string{
		"https://pkg.go.dev",
		"https://pkg.go.dev/golang.org/x",
	}

	urls := make([]*url.URL, len(urlStrings))
	parseErrs := make([]error, len(urlStrings))

	for i, urlString := range urlStrings {
		urls[i], parseErrs[i] = url.Parse(urlString)
	}

	responses, networkErrs := fetch.Fetch(urls)

	results := make([][]string, len(urlStrings))
	parseLinkErrs := make([]error, len(urlStrings))

	for i, response := range responses {
		results[i], parseLinkErrs[i] = fetch.ParseHyperLinks(response)
	}

	for i := range len(urlStrings) {
		fmt.Printf("link: %s\n", urlStrings[i])
		if parseErrs[i] != nil {
			fmt.Printf("error while checking URL validity: %+v\n", parseErrs[i])
		}
		if networkErrs[i] != nil {
			fmt.Printf("error while fetching HTTP response: %+v\n", networkErrs[i])
		}
		if parseLinkErrs[i] != nil {
			fmt.Printf("error while parsing hyperlinks in response: %+v\n", parseLinkErrs[i])
		}
		if len(results[i]) > 0 {
			fmt.Printf("hyperlinks in the page %s\n", urlStrings[i])
			for i, link := range results[i] {
				fmt.Printf("link %d %s\n", i, link)
			}
		}
	}
}
