// SILVER - Service Wrapper
//
// Copyright (c) 2016 PaperCut Software http://www.papercut.com/
// Use of this source code is governed by an MIT or GPL Version 2 license.
// See the project's LICENSE file for more information.
//

// +build ignore

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 10)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go sayHelloForever()

	<-c
	fmt.Println("Shutting down...but I'm busy and will take a while!")
	time.Sleep(1 * time.Hour)
	os.Exit(0)
}

func sayHelloForever() {
	for {
		fmt.Printf("Hello World at %v\n", time.Now())
		time.Sleep(1 * time.Second)
	}
}
