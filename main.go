package main

import (
	"fmt"
	"os"

	netcat "netcat/handler"
	//"netcat/handler/netc"
)

const usage = "[USAGE]: ./TCPCHAT $port"

func main() {
	port := netcat.Checkargs(os.Args)
	if port == "" {
		fmt.Println(usage)
		return
	}
	fmt.Printf("starting server on port %s\n", port)
	server := netcat.NewServere(port)
	server.Srart()
}
