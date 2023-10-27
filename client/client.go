package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6969")
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	go receiveMessages(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		msg := scanner.Text()
		if msg == "exit" {
			break
		}
		_, err := conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}
}

func receiveMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed by the server.")
			os.Exit(0)
		}
		fmt.Print(msg)
	}
}
