package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

var users = make(map[string]net.IPAddr)

func handler(conn net.Conn) {
	rdr := bufio.NewReader(conn)
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			log.Printf("user %v disconnected", conn.RemoteAddr())
			return
		} else if err != nil {
			fmt.Println("error while reading from client: ", err)
			return
		}
		log.Printf("input received from %v: %v\n", conn.RemoteAddr(), line)
	}
}

const ServerPort = 8080

func main() {
	ln, err := net.Listen("tcp4", ":8080")
	defer ln.Close()
	if err != nil {
		log.Fatalf("Error while trying to listen on port %v: %v\n", ServerPort, err)
	}
	log.Printf("Listening for connections on %v\n", ServerPort)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error while accepting connection: %v\n", err)
		}
		log.Printf("Connected to %v", conn.RemoteAddr().String())
		go handler(conn)
	}
}
