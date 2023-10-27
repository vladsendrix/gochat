package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	serverPtr := flag.String("s", "localhost", "Server address")
	portPtr := flag.String("p", "6969", "Port number")
	flag.Parse()

	serverAddress := *serverPtr + ":" + *portPtr
	conn, err := net.Dial("tcp", serverAddress)
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
