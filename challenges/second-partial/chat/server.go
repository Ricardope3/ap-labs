// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

type userStruct struct {
	userName string
	msg      string
}

var (
	entering     = make(chan client)
	leaving      = make(chan client)
	messages     = make(chan string) // all incoming client messages
	direct       = make(chan userStruct)
	serverPrefix = "irc-server > "
	admin        client
	globalUser   string
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	users := make(map[string]client)
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case uS := <-direct:
			users[uS.userName] <- uS.msg

		case cli := <-entering:
			if len(clients) == 0 {
				cli <- serverPrefix + "Congrats, you were the first user."
				cli <- serverPrefix + "You're the new IRC Server ADMIN"
				fmt.Printf("[%s] was promoted as the channel ADMIN\n", globalUser)
				admin = cli
			}
			clients[cli] = true
			users[globalUser] = cli

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	//GET THE USER NAME
	var buf = make([]byte, 1024)
	conn.Read(buf)
	localUser := string(bytes.Trim(buf, "\x00"))
	globalUser = string(bytes.Trim(buf, "\x00"))

	fmt.Printf("%sNew connected user [%s]\n", serverPrefix, localUser)

	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	//Welcoming messages
	ch <- serverPrefix + "Welcome to the Simple IRC Server"
	ch <- serverPrefix + "Your user [" + localUser + "] is successfully logged"

	//Entering messages
	messages <- serverPrefix + "New connected user [" + localUser + "]"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		if len(input.Text()) > 0 && string(input.Text()[0]) == "/" {
			slice := strings.Split(input.Text(), " ")
			command := slice[0]
			switch command {
			case "/users":
				messages <- "users"
			case "/msg":
				addressee := slice[1]
				directMessage := input.Text()[strings.Index(input.Text(), addressee)+len(addressee)+1:]
				direct <- userStruct{addressee, localUser + " > " + directMessage}

			case "/time":
				messages <- "tiempo"
			case "/user":
				messages <- "usuario"
			case "/kick":
				messages <- "amonos alv"

			}

		} else {
			messages <- localUser + " > " + input.Text()
		}
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- "[" + localUser + "] has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: go run server.go -host localhost -port <port>")
		os.Exit(1)
	}
	host := os.Args[2]
	port := os.Args[4]

	listener, err := net.Listen("tcp", host+":"+port)
	fmt.Println(serverPrefix + "Simple IRC Server started at " + host + ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	fmt.Println(serverPrefix + "Ready for receiving new clients")
	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
