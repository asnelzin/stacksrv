package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func Server(port int, stack *Stack) {
	log.Printf("[INFO] activate tcp server on port %d", port)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("[ERROR] could not create listener: %v", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("[ERROR] could not accept new connection: %v", err)
		}

		go handleRequest(conn, stack)
	}
}

func handleRequest(conn net.Conn, stack *Stack) {
	log.Print("[INFO] got a new connection")

	r := bufio.NewReader(conn)
	header, err := r.ReadByte()
	if err != nil {
		log.Printf("[ERROR] could not read header from connection: %v", err)
	}

	op := header >> 7
	switch op {
	case 0:
		n, err := conn.Write(stack.Pop())
		if err != nil {
			log.Printf("[ERROR] could not write response to connection: %v", err)
		}
		log.Printf("[DEBUG] write %d bytes", n)
	case 1:
		length := header & 0x7f
		buf := make([]byte, length)

		r.Read(buf)
		stack.Push(buf)
	}
	conn.Close()
}
