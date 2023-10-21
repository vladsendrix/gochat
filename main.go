package main

import (
	"fmt"
	"github.com/vladsendrix/gochat/server"
)

func main() {
	fmt.Println("Welcome to Gochat")
	server.Start("localhost:6969")
}
