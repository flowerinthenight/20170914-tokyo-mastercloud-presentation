package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func signalHandler() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(
		sigchan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		for {
			s := <-sigchan
			switch s {
			case syscall.SIGINT, syscall.SIGTERM:
				log.Println("Terminated.")
				os.Exit(0)
			}
		}
	}()
}

func main() {
	signalHandler()
	log.SetFlags(0)

	for {
		log.Println("time now:", time.Now())
		time.Sleep(time.Second * 1)
	}
}
