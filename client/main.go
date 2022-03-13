package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func errorLog(err error, name string) {
	if err != nil {
		log.Fatalf("%v function failed to execute", name)
	}
}
func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	errorLog(err, "dial")

	defer connection.Close()
	//&{{0xc00013a080}}
	fmt.Println("Enter your username: ")

	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	errorLog(err, "reading")

	username = strings.Trim(username, " \r\n")

	welcomeMsg := fmt.Sprintf("welcome %s, to the chat, say hi to your village people.\n", username)
	fmt.Println(welcomeMsg)
	//connection.Write([]byte(welcomeMsg))
	go read(connection)
	write(connection, username)

	// read
	// write
}
func read(connection net.Conn) {
	for {
		reader := bufio.NewReader(connection)
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			connection.Close()
			fmt.Println("connection closed")
			os.Exit(0)
		}
		fmt.Println(message)
		fmt.Println("------------------------------------------")
	}

}
func write(connection net.Conn, username string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		messagge, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// formats the string to -> username: - message
		messagge = fmt.Sprintf("%s:-%s\n", username, strings.Trim(messagge, " \r\n"))

		connection.Write([]byte(messagge))
	}
}
