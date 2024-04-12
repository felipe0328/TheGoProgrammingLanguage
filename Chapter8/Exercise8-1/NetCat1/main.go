package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var port = flag.Int("port", 8080, "port to use as clock")

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
