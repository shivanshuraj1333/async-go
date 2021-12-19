package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listner, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			return
		}
	}(c)
	for {
		_, err := io.WriteString(c, "respose from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
