package main

import (
	"fmt"
	"os"
	"strconv"
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
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	f, err := os.Create("ping-pong.txt")
	check(err)
	outputString :=
		`THIS IS THE OUTPUT FILE FOR PING-PONG.GO
----------------------------------------
- I made two unbufered chanels, PING and PONG
- I first run a goroutine that sends a message to PING, this kickstarts the whole process
- Afterwardes there's an infinite for loop with a SELECT statement
- The messages go back and forth
- The number of communications per seconds is roughly: `
	wrote := false
	n3, err := f.WriteString(outputString)
	check(err)
	if n3 == -1 {
		fmt.Println("Error")
	}

	ping = make(chan string)
	pong = make(chan string)
	n := 0
	start := time.Now()
	go func() {
		ping <- "PING"
	}()

	for {
		select {
		case msg := <-pong:
			n++
			fmt.Println("pong recivio ", msg)
			go func() {
				n++
				fmt.Println("PONG ENVIA")
				ping <- "PING"
			}()
		case msg := <-ping:
			n++
			fmt.Println("ping recivio", msg)
			go func() {
				n++
				fmt.Println("PING ENVIA")
				pong <- "PONG"
			}()
		default:
			elapsed := time.Since(start)
			seconds := elapsed.Seconds()
			if !wrote && seconds > 1.0 && seconds < 1.001 {
				f.WriteString(strconv.Itoa(n))
				wrote = true
			}

		}
	}
}

//!-
