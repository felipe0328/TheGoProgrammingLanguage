// Modify clock2 to accept a port number, and write a program, clockwall,
// that acts as a client of several clock servers at once, reading the times
// from each one and dsiplaying the results in a table, akin to the wall of clocks
// seen in some business offices.
// TZ=US/Eastern go run ./Clock -port 8010 &
// TZ=Asia/Tokio go run ./Clock -port 8020 &
// TZ=Europe/London go run ./Clock -port 8030
// clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	inputs := os.Args[1:]

	for _, input := range inputs {
		data := strings.Split(input, "=")
		location, port := data[0], data[1]
		go connectAndGetTime(location, port)
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

func connectAndGetTime(location, port string) {
	fmt.Println(location)
	conn, err := net.Dial("tcp", port)
	if err != nil {
		fmt.Printf("Error Connect to %s: %s\n", location, err)
		return
	}

	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
