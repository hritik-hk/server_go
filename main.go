package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const threadPool_Size = 100

func handleConnection(conn net.Conn) {
	requestBuffer := make([]byte, 2048)

	// Read the HTTP request
	n, err := conn.Read(requestBuffer)
	if err != nil {
		log.Printf("Error reading request: %v\n", err)
		conn.Close()
		return
	}

	request := string(requestBuffer[:n])

	client := strings.Split(request, "\n")[7:8]

	// Mimicking Long-running Job
	log.Printf("processing the request... client- %v ", client[0])
	time.Sleep(5 * time.Second)

	// Returning the response and closing
	log.Printf("processing complete... client- %v ", client[0])
	resp := fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\nHello, client %v \r\n", client[0])
	conn.Write([]byte(resp))
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

	pool := threadPool(threadPool_Size)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		log.Printf("new client connected")

		go handleConnection(conn)

		job := func() {
			handleConnection(conn)
		}
		pool.AddJob(job)

	}
}
