package main

import (
	"flag"
	"log"
	"time"
)

// func pingServer(pingHost string, pingPort string) error {
// 	log.Println("before")
// 	conn, err := net.DialTCP(net.JoinHostPort(pingHost, pingPort))
// 	defer conn.Close()
// 	log.Println("after")
// 	if err != nil {
// 		return err
// 	}
//
// 	for {
// 		// _, err := conn.Write([]byte(":ping:"))
// 		// if err != nil {
// 		// 	log.Println(err)
// 		// 	ch <- err
// 		// }
// 		log.Println("connection ok")
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	var user string
	var serverHost string
	//var dataPort string
	var serverPort string

	flag.StringVar(&user, "u", "", "connect to server as this user")
	//password := flag.String("p", "", "user password for authentication")
	flag.StringVar(&serverHost, "s", "127.0.0.1", "server name/address")
	flag.StringVar(&serverPort, "p", "8080", "ping port")
	flag.Parse()

	cl, err := newClient(user, serverHost, serverPort)
	if err != nil {
		log.Fatalln(err)
	}

	defer cl.close()

	cl.send("hello")
	time.Sleep(time.Second * 10)
	cl.send("hi")
	time.Sleep(time.Second * 10)
	cl.send("bye")
	time.Sleep(time.Second * 10)
}
