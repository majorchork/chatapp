package main

import (
	"bufio"
	"fmt"
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
	connection := <-newConnection
	reader := bufio.NewReader(connection)
	message, err := reader.ReadString('\n')
	errorLog(err, "reader")
	fmt.Println(message)

}
