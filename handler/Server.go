package natc

import (
	"bufio"
	"sync"
)

type Server struct {
	port string

	clinte   map[*clinte]bool
	messages chan message
	history  []string
	mux      *sync.Mutex
	// server.mux.Lock()
	// server.clients[newClient] = true
	// server.mux.Unlock()
}

type clinte struct {
	Name   string
	Writer *bufio.Writer
}

type message struct {
	message string
	sender  *clinte
}

func NewServere(port string) *Server {
	return &Server{
		port:     port,
		clinte:   make(map[*clinte]bool),
		messages: make(chan message),
		history:  []string{},
	}
}
