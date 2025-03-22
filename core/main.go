package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MemKV server running on port 8081...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn, "Welcome to MemKV!\n")
}
