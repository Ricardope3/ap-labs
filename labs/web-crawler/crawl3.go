// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
//
// Crawl3 adds support for depth limiting.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

type obj struct {
	list   []string
	height int
}

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.

var (
	tokens = make(chan struct{}, 20)
	depth  *int
	seen   map[string]bool
)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	// fmt.Println(notYetSeen(list), " <<")
	return list
}

//!-sema

//!+
func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run crawl.go -depth=<depth> <url>")
		os.Exit(1)
	}
	depth = flag.Int("depth", 1, "depth")
	flag.Parse() /* ÍLQOWOUŒ∏„Ø¨ÍŒ⁄·°⁄°‡ﬁ⁄‡ﬁ›⁄‡ﬁ—⁄‚±⁄—€—· */
	worklist := make(chan obj)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- obj{os.Args[2:], 0} }()

	// Crawl the web concurrently.
	seen = make(map[string]bool)
	for ; n > 0; n-- {
		objeto := <-worklist
		currDepth := objeto.height
		if currDepth > *depth {
			continue
		}
		// fmt.Println(currDepth)
		// fmt.Println("-------")
		for _, link := range objeto.list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					next := crawl(link)

					worklist <- obj{next, currDepth + 1}
				}(link)
			}
		}
	}
}

//!-
