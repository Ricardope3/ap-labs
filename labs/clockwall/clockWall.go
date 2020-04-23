package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func handleConection(location, host string, wait chan int) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan int)
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- 2 // signal the main goroutine
		wait <- 1
	}()
	x := 1
	x = <-done // wait for background goroutine to finish
	log.Println("Channel Closed with value: ", x)
	close(done)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("go run clockWall.go <location>=localhost:<port> <location>=localhost:<port> ...")
		os.Exit(1)
	}

	argumentos := os.Args[1:]

	wait := make(chan int)

	for _, val := range argumentos {
		array := strings.Split(val, "=")
		location := array[0]
		host := array[1]

		go handleConection(location, host, wait)
	}
	<-wait
	close(wait)
}
