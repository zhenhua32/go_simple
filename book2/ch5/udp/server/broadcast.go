package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type connections struct {
	addrs map[string]*net.UDPAddr
	mu    sync.Mutex
}

func broadcast(conn *net.UDPConn, conns *connections) {
	count := 0
	for {
		conns.mu.Lock()
		for _, retAddr := range conns.addrs {
			fmt.Println("send msg to ", retAddr)
			msg := fmt.Sprintf("Sent %d", count)
			if _, err := conn.WriteToUDP([]byte(msg), retAddr); err != nil {
				fmt.Printf("error encountered: %s", err.Error())
				continue
			}
		}

		conns.mu.Unlock()
		time.Sleep(1 * time.Second)
		count++
	}
}
