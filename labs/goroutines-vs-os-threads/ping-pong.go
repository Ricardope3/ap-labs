package main

import (
	"fmt"
	"time"
)

var (
	ping chan string
	pong chan string
)

func sendToPing() {
	fmt.Println("Sending to ping")
	ping <- <-pong
}

func sendToPong() {
	fmt.Println("Sending to pong")
	pong <- <-ping
}

func main() {
	start := time.Now()
	ping = make(chan string)
	pong = make(chan string)
	n := 0
	go func() {
		ping <- "PING"
	}()
	go func() {
		pong <- "PONG"
	}()

	for {
		select {
		case <-pong:
			n++
			// fmt.Println("pong recivio ", msg)
			go func() {
				// fmt.Println("PONG ENVIA")
				ping <- "PING"
			}()
		case <-ping:
			n++
			// fmt.Println("ping recivio", msg)
			go func() {
				// fmt.Println("PING ENVIA")
				pong <- "PONG"
			}()
		default:
			n++
			elapsed := time.Since(start)
			seconds := elapsed.Seconds()
			if seconds > 1.0 && seconds < 1.0001 {
				fmt.Println(n)
			}

		}
	}
}

//!-
