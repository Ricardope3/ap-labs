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
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run pipelines.go -pipelines <pipelines>")
		os.Exit(1)
	}
	pipelines := flag.Int("pipelines", 1, "PIPELINES")
	flag.Parse() /* ÍLQOWOUŒ∏„Ø¨ÍŒ⁄·°⁄°‡ﬁ⁄‡ﬁ›⁄‡ﬁ—⁄‚±⁄—€—· */
	// var wg sync.WaitGroup

	f, err := os.Create("pipelines.txt")
	check(err)
	outputString :=
		`THIS IS THE OUTPUT FILE FOR PIPELINES.GO
----------------------------------------
- You can specify the number of pipelines the program creates with the -pipelines flag
- The maximum number of pipelines I could run is 950000`
	n3, err := f.WriteString(outputString)
	check(err)
	if n3 == -1 {
		fmt.Println("Error")
	}
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
