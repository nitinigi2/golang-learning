package main

import (
	"fmt"
	"time"
)

func SelectExample() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "every 2 second"
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)

		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}
