package main

import (
	"bufio"
	"log"
	"net"
)

func errorLog(err error, name string) {
	if err != nil {
		log.Fatalf("%v function failed to execute", name)
	}
}

var (
	openConnections = make(map[net.Conn]bool)
	newConnection   = make(chan net.Conn)
	deadConnection  = make(chan net.Conn)
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	errorLog(err, "listen")

	defer ln.Close()
	go func() {
		for {
			conn, err := ln.Accept()
			errorLog(err, "accept")

			openConnections[conn] = true
			newConnection <- conn
		}
	}()
	//connection := <-newConnection
	//reader := bufio.NewReader(connection)
	//message, err := reader.ReadString('\n')
	//errorLog(err, "reader")
	//fmt.Println(message)
	for {
		select {
		case conn := <-newConnection:
			go broadcastMessage(conn)
			// invoke broadcast(broadcast's to other connections)
		case conn := <-deadConnection:
			for item := range openConnections {
				if item == conn {
					break
				}
			}
			delete(openConnections, conn)
		}
	}
}
func broadcastMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)

		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		for item := range openConnections {
			if item != conn {
				item.Write([]byte(message))
			}
		}
	}
	deadConnection <- conn
}
