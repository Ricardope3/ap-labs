package main

import (
	"flag"
	"fmt"
	"os"
)

var wait chan int

func doSomething(primero, segundo chan int) {
	var x int = <-primero
	segundo <- x

}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run pipelines.go -pipelines <pipelines>")
		os.Exit(1)
	}
	pipelines := flag.Int("pipelines", 1, "PIPELINES")
	flag.Parse() /* ÍLQOWOUŒ∏„Ø¨ÍŒ⁄·°⁄°‡ﬁ⁄‡ﬁ›⁄‡ﬁ—⁄‚±⁄—€—· */
	// var wg sync.WaitGroup
	array := make([]chan int, *pipelines)
	wait = make(chan int)
	for i := 0; i < *pipelines; i++ {
		array[i] = make(chan int)
	}

	go func() {
		array[0] <- 290
	}()

	for i := 0; i < *pipelines-1; i++ {
		go doSomething(array[i], array[i+1])

	}

	fmt.Println(<-array[*pipelines-1])
}

//!-
