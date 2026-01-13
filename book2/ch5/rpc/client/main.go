package main

import (
	"book2/ch5/rpc/tweak"
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("error dialing:", err)
	}

	args := tweak.Args{
		String:  "hello, world",
		ToUpper: true,
		Reverse: true,
	}

	var result string
	err = client.Call("StringTweak.Tweak", args, &result)
	if err != nil {
		log.Fatal("client call with error:", err)
	}
	fmt.Printf("the result is: %s", result)
}
