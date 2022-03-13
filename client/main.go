package main

import (
	"bufio"
	"fmt"
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
	m := []byte(welcomeMsg)
	connection.Write(m)
}
