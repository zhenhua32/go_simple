package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func CatchSig(ch chan os.Signal, done chan bool) {
	sig := <-ch
	fmt.Println("nsig received:", sig)

	switch sig {
	case syscall.SIGINT:
		fmt.Println("handing a SIGINT now")
	case syscall.SIGTERM:
		fmt.Println("handing a SIGTERM now")
	default:
		fmt.Println("unknown signal:", sig)
	}

	done <- true
}

func main() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go CatchSig(signals, done)
	fmt.Println("press ctrl+c to terminate...")
	<-done
	fmt.Println("Done!")
}
