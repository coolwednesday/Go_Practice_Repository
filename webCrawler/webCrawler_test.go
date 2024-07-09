package webCrawler

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func compareUrls(i, j int) bool {
	return s[i].url < s[j].url
}

func TestWebCrawler(t *testing.T) {

	validOutput1 := []res{
		{url: "https://golang.org/", body: "The Go Programming Language", found: true},
		{url: "https://golang.org/cmd/", body: "", found: false},
		{url: "https://golang.org/pkg/", body: "Packages", found: true},
		{url: "https://golang.org/pkg/fmt/", body: "Package fmt", found: true},
		{url: "https://golang.org/pkg/os/", body: "Package os", found: true},
	}

	tests := []struct {
		name   string
		input  string
		output []res
	}{
		{name: "Valid URL", input: "https://golang.org/", output: validOutput1},
		{name: "Valid URL", input: "https://golang.org/pkg/", output: validOutput1},
		{name: "Invalid URL", input: "http://www.google.com", output: []res{{url: "http://www.google.com", body: "", found: false}}},
		{name: "Empty URL", input: "", output: []res{{url: "", body: "", found: false}}},
	}

	for _, test := range tests {
		o := fetchUrls(test.input)
		//sorting the slice
		sort.Slice(s, compareUrls)
		assert.Equal(t, test.output, o)
		s = nil
	}
}
