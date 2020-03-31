package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var mapa = map[string]int{}
	for _, word := range strings.Split(string(s), " ") {
		mapa[word]++
	}

	return mapa
}

func main() {
	wc.Test(WordCount)
}
