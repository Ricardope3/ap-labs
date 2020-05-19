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

	"gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
var seen = make(map[string]bool)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema
func crawler(url string, depth int, done chan<- bool) {
	defer func() { done <- true }()

	if depth < 0 {
		return
	}

	_, ok := seen[url]
	if ok {
		return
	}
	seen[url] = true

	list := crawl(url)
	for _, link := range list {
		recursion := make(chan bool)
		go crawler(link, depth-1, recursion)
		<-recursion
	}
}

//!+
func main() {
	worklist := make(chan []string)
	done := make(chan bool)

	var depth = flag.Int("depth", 3, "Depth of crawl")
	flag.Parse()
	go func() { worklist <- flag.Args() }()

	url := <-worklist

	go crawler(url[0], *depth, done)
	<-done
}

//!-
