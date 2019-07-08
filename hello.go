// all programs start running in package main
package main

import (
    "fmt"
    "sync"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

type SafeMap struct {
    v map[string]string
    mux sync.Mutex
}

// map with URL string keys for string bodies
var crawled = SafeMap{v: make(map[string]string)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth
func Crawl(url string, depth int, fetcher Fetcher) { 
    // TODO fetch URLs in parallel
    // TODO don't fetch the same URL twice
    // This implementation doesn't do either

    if depth <= 0 {
        return
    }

    // check if we've already been to this URL, leave if so
    crawled.mux.Lock()
    _, visited := crawled.v[url]

    if visited {
        crawled.mux.Unlock()
        return
    } else {
        crawled.v[url] = "body pending"
        crawled.mux.Unlock()
    }

    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }

    crawled.v[url] = body

    fmt.Printf("found: %s %q\n", url, body)

    // channel from golang solution
    done := make(chan bool)

    for _, u := range urls {
            go func(url string) {
                Crawl(url, depth-1, fetcher)
                done <- true
            }(u)
    }

    // wait for everything to return
    for range urls {
        <- done
    }

    return
}

func main() {
    Crawl("https://golang.org/", 4, fetcher)

}

// fakeFetcher is a fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher
var fetcher = fakeFetcher {
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
