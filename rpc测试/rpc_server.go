package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello, " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	log.Println("RPC server started on :1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
			continue
		}
		rpc.ServeConn(conn)
	}
}
