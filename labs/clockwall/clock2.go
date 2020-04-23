// Clock2 is a concurrent TCP server that periodically writes the time.
package main

import (
	// "io"
	// "log"
	// "net"
	// "time"
	"os"
	"fmt"
)

// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
// 		if err != nil {
// 			return // e.g., client disconnected
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }

func main() {
	tz := os.Getenv("TZ")
	if len(os.Args) < 2 || len(tz) <= 0 {
		fmt.Println("Usage: TZ=<timezone> go run clock2.go -port <port>")
		os.Exit(1)
	}
	port := os.Args[2]
	fmt.Println(tz)
	fmt.Println(port)
	
	// listener, err := net.Listen("tcp", "localhost:9090")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		log.Print(err) // e.g., connection aborted
	// 		continue
	// 	}
	// 	go handleConn(conn) // handle connections concurrently
	// }
}
