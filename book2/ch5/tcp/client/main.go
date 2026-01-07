package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const addr = "localhost:8080"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("enter some text:")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("reading input error:%s\n", err.Error())
			continue
		}

		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("error connecting to server:%s\n", err.Error())
			continue
		}

		fmt.Fprint(conn, data)

		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("error reading from server:%s\n", err.Error())
			continue
		}

		fmt.Printf("received from server:%s", status)

		conn.Close()
	}
}
