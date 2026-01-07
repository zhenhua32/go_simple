package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const addr = "localhost:8080"

func echoBackCapitalized(conn net.Conn) {
	reader := bufio.NewReader(conn)

	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading data:%s\n", err.Error())
		return
	}

	fmt.Printf("received: %s", data)
	conn.Write([]byte(strings.ToUpper(data)))

	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	fmt.Printf("server listening on %s\n", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("error accepting connection:%s\n", err.Error())
			continue
		}

		go echoBackCapitalized(conn)
	}
}
