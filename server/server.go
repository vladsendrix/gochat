package server

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

type ChatServer struct {
	clients map[net.Conn]struct{}
	mutex   sync.RWMutex
}

func Start(address string) {
	server := NewChatServer()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Chat server started on %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go server.HandleClient(conn)
	}
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		clients: make(map[net.Conn]struct{}),
	}
}

func (s *ChatServer) HandleClient(conn net.Conn) {
	defer conn.Close()

	s.mutex.Lock()
	s.clients[conn] = struct{}{}
	s.mutex.Unlock()

	defer func() {
		s.mutex.Lock()
		delete(s.clients, conn)
		s.mutex.Unlock()
	}()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		msg := strings.TrimSpace(string(buf[:n]))

		if msg != "" {
			fmt.Println("Received:", msg)
			s.Broadcast(msg)
		}
	}
}

func (s *ChatServer) Broadcast(msg string) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for conn := range s.clients {
		_, err := conn.Write([]byte(msg + "\n"))
		if err != nil {
			delete(s.clients, conn)
		}
	}
}
