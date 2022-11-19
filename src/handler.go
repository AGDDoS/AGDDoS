package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func SetupCloseHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("\r - Target URL: " + TargetUrl)
		log.Println("\r - Duration Minute: ", DurationMinute)
		log.Println("\r - Total Requests: ", Totalrequest)
		os.Exit(0)
	}()
}
