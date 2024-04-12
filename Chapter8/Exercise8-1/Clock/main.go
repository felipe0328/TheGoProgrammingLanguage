package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 8080, "port to use as clock")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // eg. Connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		currTime := fmt.Sprintf("%s\n", time.Now().Format("15:04:05"))
		_, err := io.WriteString(c, currTime)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
