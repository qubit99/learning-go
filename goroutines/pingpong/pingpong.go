package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const address = "localhost:8080"

func runServer() error {
	conn, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error runninng server:", err)
		return err
	}
	fmt.Println("Server is running on ", address)

	defer conn.Close()

	for {
		client, err := conn.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConn(client)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
		message = strings.TrimSpace(message)
		fmt.Println("Message received from client: ", message)
		if strings.EqualFold(message, "PING") {
			go func() {
				time.Sleep(200 * time.Millisecond)
				c.Write([]byte("PONG\n"))
			}()
		} else {
			go func() { _, _ = c.Write([]byte("unknown\n")) }()
		}
	}
}

func runClient() error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return err
	}
	defer conn.Close()

	go func() {
		sc := bufio.NewScanner(conn)
		for sc.Scan() {
			fmt.Println("client recv:", sc.Text())
		}
		if err := sc.Err(); err != nil {
			fmt.Println("client read error:", err)
		}
	}()

	_, err = conn.Write([]byte("ping\n"))
	if err != nil {
		return err
	}

	// keep client alive long enough to see the async response
	time.Sleep(10 * time.Second)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run pingpong.go [server|client]")
		return
	}
	switch os.Args[1] {
	case "server":
		if err := runServer(); err != nil {
			fmt.Println("server error:", err)
		}
	case "client":
		if err := runClient(); err != nil {
			fmt.Println("client error:", err)
		}
	default:
		fmt.Println("usage: go run pingpong.go [server|client]")
	}
}
