package webCrawler

import (
	"fmt"
	"sync"
)

type res struct {
	url   string
	body  string
	found bool
}

var s []res

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func CrawlHelper(url string, depth int, fetcher Fetcher) {
	fetched := make(map[string]bool)
	Crawl(url, depth, fetcher, fetched)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.

func Crawl(url string, depth int, fetcher Fetcher, fetched map[string]bool) {
	// We have to Fetch URLs in parallel.
	// Here we do not have to fetch the same URL twice.
	if depth <= 0 {
		return
	}
	if fetched[url] == true {
		return
	} else {
		fetched[url] = true
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		s = append(s, res{url: url, body: body, found: false})
		fmt.Println(err)
		return
	}
	c := res{
		url:   url,
		body:  body,
		found: true,
	}
	s = append(s, c)
	fmt.Printf("found: %s %q\n", url, body)

	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			Crawl(u, depth-1, fetcher, fetched)
			wg.Done()
		}(u)
	}
	wg.Wait()
	return
}

func fetchUrls(url string) []res {
	CrawlHelper(url, 4, fetcher)
	return s
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
