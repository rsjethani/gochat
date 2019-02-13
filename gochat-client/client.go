package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	user   string
	server net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

func (cl *client) send(msg string) error {
	_, err := cl.writer.WriteString(fmt.Sprint(msg, "\n"))
	if err != nil {
		return fmt.Errorf("Error while sending data to server: %v", err)
	}

	err = cl.writer.Flush()
	if err != nil {
		return fmt.Errorf("Error while flushing write buffer: %v", err)
	}

	return nil
}

func (cl *client) login() error {
	msg := fmt.Sprintf(":logirggggggggggggggggggggggggggggggggsgggrhttyin: %v", cl.user)
	err := cl.send(msg)
	if err != nil {
		return fmt.Errorf("Error while logging into the server: ", err)
	}
	return nil
}

func newClient(user string, serverHost string, serverPort string) (*client, error) {
	serverAddr, err := net.ResolveTCPAddr("tcp4", net.JoinHostPort(serverHost, serverPort))
	if err != nil {
		return nil, fmt.Errorf("Error while resolving Server address %v: %v", serverAddr, err)
	}
	conn, err := net.DialTCP("tcp4", nil, serverAddr)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting to server at %v: %v", serverAddr, err)
	}
	log.Printf("Connection to server %v successful", serverAddr)

	cl := client{
		user:   user,
		server: conn,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}

	err = cl.login()
	if err != nil {
		return nil, fmt.Errorf("Error while logging into server %v: %v", serverAddr, err)
	}

	return &cl, nil
}

func (cl *client) close() {
	err := cl.server.Close()
	if err != nil {
		log.Println("Error while closing client connection: ", err)
	}
}
