package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time" // or "runtime"
)

func cleanup() {
	fmt.Println("cleanup")
}

func main() {
	var option int
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	for {
		fmt.Println("sleeping...")
		fmt.Scanf("%d\n",&option)
		time.Sleep(2 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}
}
