package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

func EchoServer(ws *websocket.Conn) {
	file, _ := os.ReadFile("a.txt")
	n, err := ws.Write(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n, "bytes written")
}

func main() {
	http.Handle("/ws", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
