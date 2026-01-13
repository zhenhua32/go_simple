package main

import (
	"book2/ch5/rpc/tweak"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	s := new(tweak.StringTweak)
	if err := rpc.Register(s); err != nil {
		log.Fatal("failed to register: ", err)
	}

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen err:", err)
	}

	fmt.Println("listening on :1234")
	log.Panic(http.Serve(l, nil))
}
