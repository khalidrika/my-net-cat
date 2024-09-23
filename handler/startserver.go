package natc

import (
	"fmt"
	"net"
	"os"
)

func (s *Server) Srart() {
	listener, err := net.Listen("tcp", ":"+s.port)
	defer listener.Close()
	if err != nil {
		fmt.Println("error with accepting")
		return
	}
	go s.brc()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error with accepting")
			return
		}
		go s.handelclint(conn)
	}
}

func (s *Server) brc() {
	for {
		msg := <-s.messages
		for clint := range s.clinte {
			if msg.sender != clint {
				_, err := clint.Writer.WriteString(msg.message)
				if err != nil {
					fmt.Println("error broadcasting")
					os.Exit(1)
				}
				err = clint.Writer.Flush()
				if err != nil {
					fmt.Println("error flushing")
					os.Exit(1)
				}
				sendPrompt(clint)
			}
		}
	}
}

func (s *Server) handelclint(conn net.Conn) {
	defer conn.Close()
	if len(s.clinte) == 10 {
		conn.Write([]byte("maximum 10 connectinons availble\n"))
		return
	}

	WriteWelcome(conn)
	namebuffer, err := readName(conn)
	if err != nil {
		fmt.Println("error reading a name")
		return
	}
	cleint := newClient(namebuffer, conn)
	s.joinchat(cleint)
}
func (s *Server) joinchat(clint *clinte) {
	s.addclinte(clint)
	msg := "\n" + clint.Name + " has join the chat .\n"
	s.messages <- message{msg, clint}

	s.showhistory(clint)
	s.addtohistory(msg)
}
func (s *Server) addclinte(clint *clinte) {
	s.mux.Lock()
	s.clinte[clint] = true
	s.mux.Unlock()
}

func (s -*Server) showhistory(clint *clinte) {
	for _, msg := range s.history {
		msg
	}
}