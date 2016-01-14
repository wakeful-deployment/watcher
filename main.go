package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGHUP)

	go RespondToHub(signals)

	for {
	}
}

func RespondToHub(signals chan os.Signal) {
	for range signals {
		fmt.Println("Got a SIGHUP! Now restarting haproxy...")
		RestartHAProxy()
	}
}

func RestartHAProxy() {
	fmt.Println("Restarting HAProxy")
}
