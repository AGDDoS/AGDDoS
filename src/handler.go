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
		log.Println("\r --- AGDDoS Report for " + TargetUrl)
		// log.Println("\r - DurationMinute(dm):", DurationMinute)
		log.Println("\r - Total Requests:", Totalrequest)
		os.Exit(0)
	}()
}
