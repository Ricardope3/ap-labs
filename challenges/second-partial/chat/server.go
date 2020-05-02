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
	"time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

type userStruct struct {
	userName string
	msg      string
}

type clientStruct struct {
	cliente    client
	connection net.Conn
}

// type generalStruct struct {
// 	cliente    client
// 	connection net.Conn
// 	user       string
// }

var (
	serverPrefix = "irc-server > "
	admin        client
	globalUser   string
	entering     = make(chan clientStruct)
	leaving      = make(chan client)
	messages     = make(chan string) // all incoming client messages
	direct       = make(chan userStruct)
	kick         = make(chan string)
	users        map[string]client
	connections  map[string]net.Conn
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	users = make(map[string]client)
	connections = make(map[string]net.Conn)

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

		case clientStructure := <-entering:
			if len(clients) == 0 {
				clientStructure.cliente <- serverPrefix + "Congrats, you were the first user."
				clientStructure.cliente <- serverPrefix + "You're the new IRC Server ADMIN"
				fmt.Printf("[%s] was promoted as the channel ADMIN\n", globalUser)
				admin = clientStructure.cliente
			}
			clients[clientStructure.cliente] = true
			users[globalUser] = clientStructure.cliente
			connections[globalUser] = clientStructure.connection

		case cli := <-leaving:
			if admin == cli {

				for newAdmin := range clients {
					admin = newAdmin
					newAdmin <- serverPrefix + "You're the new admin!"
					continue
				}
			}

			delete(clients, cli)
			close(cli)
		case user := <-kick:
			connection := connections[user]
			client := users[user]
			delete(clients, client)
			delete(connections, user)
			delete(users, user)
			connection.Close()
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

	ch := make(chan string) // outgoing client messages

	if users[localUser] != nil {
		fmt.Fprintln(conn, "User with name: "+localUser+" already exists")
		close(ch)
		conn.Close()
		return
	}

	fmt.Printf("%sNew connected user [%s]\n", serverPrefix, localUser)

	go clientWriter(conn, ch)

	//Welcoming messages
	ch <- serverPrefix + "Welcome to the Simple IRC Server"
	ch <- serverPrefix + "Your user [" + localUser + "] is successfully logged"

	//Entering messages
	messages <- serverPrefix + "New connected user [" + localUser + "]"
	entering <- clientStruct{ch, conn}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		if len(input.Text()) > 0 && string(input.Text()[0]) == "/" {
			slice := strings.Split(input.Text(), " ")
			command := slice[0]
			switch command {
			case "/users":
				str := ""
				for usuario := range users {
					str += usuario + ", "
				}
				ch <- str
			case "/msg":
				if len(slice) < 2 {
					ch <- "Please specify a user"
					continue
				}
				if len(slice) < 3 {
					ch <- "Please enter a message"
					continue
				}
				addressee := slice[1]
				if _, ok := users[addressee]; ok {
					directMessage := input.Text()[strings.Index(input.Text(), addressee)+len(addressee)+1:]
					direct <- userStruct{addressee, localUser + " > " + directMessage}
				} else {
					ch <- "User: " + addressee + " doesn't exist"
				}

			case "/time":
				timezone := "America/Mexico_City"
				loc, _ := time.LoadLocation(timezone)
				theTime := time.Now().In(loc).Format("15:04\n")
				ch <- "Local Time: " + timezone + " " + strings.Trim(theTime, "\n")
			case "/user":
				if len(slice) < 2 {
					ch <- "Please enter a user"
					continue
				}
				user := slice[1]
				if _, ok := users[user]; ok {
					ip := connections[user].RemoteAddr().String()
					ch <- "Username: " + user + " IP: " + ip

				} else {
					ch <- "User: " + user + " doesn't exist"
				}
			case "/kick":
				if len(slice) < 2 {
					ch <- "Please enter a user to kick"
					continue
				}
				if ch == admin {
					user := slice[1]
					if _, ok := users[user]; ok {
						messages <- "[" + user + "] was kicked from channel"
						kick <- user
					} else {
						ch <- "User: " + user + " doesn't exist"
					}
				} else {
					ch <- "Only the admin can kick people out of the server"
				}
			default:
				ch <- "Invalid command"

			}

		} else {
			messages <- localUser + " > " + input.Text()
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	leaving <- ch
	messages <- "[" + localUser + "] has left"
	delete(users, localUser)
	delete(connections, localUser)
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
