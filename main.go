package main

import (
	"fmt"
	"github.com/vladsendrix/gochat/server"
)

func main() {
	fmt.Println("Welcome to Gochat")
	server.Start("0.0.0.0:6969")
}
