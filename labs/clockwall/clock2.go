// Clock2 is a concurrent TCP server that periodically writes the time.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn, timezone string) {
	defer c.Close()
	loc, _ := time.LoadLocation(timezone)
	for {
		_, err := io.WriteString(c, timezone+"\t: "+time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	timezone := os.Getenv("TZ")
	if len(os.Args) < 2 || len(timezone) <= 0 {
		fmt.Println("Usage: TZ=<timezone> go run clock2.go -port <port>")
		os.Exit(1)
	}
	port := os.Args[2]

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, timezone) // handle connections concurrently
	}
}
