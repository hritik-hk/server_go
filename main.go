package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	request := make([]byte, 1024)

	// Read the HTTP request
	conn.Read(request)

	// Mimicking Long-running Job
	log.Println("processing the request...")
	time.Sleep(5 * time.Second)

	// Returning the response and closing
	log.Println("processing complete...")
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	conn.Close()
}

func main() {

	PORT := 8080
	portString := fmt.Sprintf(":%v", PORT)

	listener, err := net.Listen("tcp", portString)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer listener.Close()
	log.Printf("Listening on PORT : %v", PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		log.Println("new client connected")

		go do(conn)
	}
}
