package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func process(c *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter some text:")
		data, err := reader.ReadString('\n')
		if err != nil {
			log.Println("read from stdin error:", err)
		}

		data = strings.TrimSpace(data)

		err = c.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println("failed to write message:", err)
			return
		}

		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("failed to read:", err)
			return
		}
		log.Printf("recv: %s", string(message))
	}
}
