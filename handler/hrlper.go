package natc

import (
	"fmt"
	"net"
	"os"
	"time"
)

func sendPrompt(clint *clinte) {
	timestamp := time.Now().Format("24-Sep-23 16:01:04 MST")
	_, err := clint.Writer.WriteString(fmt.Sprintf("[%s][%s]:", timestamp, clint.Name))
	if err != nil {
		fmt.Println("error writing string")
		os.Exit(1)
	}
	err = clint.Writer.Flush()
	if err != nil {
		fmt.Println("error flushing")
		os.Exit(1)
	}
}

func WriteWelcome(conn net.Conn) {
	var welcome string = `
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    '.       | '' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     '-'       '--'
`
	_, err := conn.Write([]byte(welcome))
	if err != nil {
		fmt.Println("error writing messagee")
		os.Exit(1)
	}
	_, err = conn.Write([]byte("enter your name : "))
	if err != nil {
		fmt.Println("error writing message name")
		os.Exit(1)
	}
}

func readName(conn net.Conn) ([]byte, error) {
	name := make([]byte, 1000)
	length, err := conn.Read(name)
	name = name[:length]
	if err != nil {
		return nil, err
	}
	return name, nil
}
