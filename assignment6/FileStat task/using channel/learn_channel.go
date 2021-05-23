package main

import (
	"fmt"
	"sync"
)

func testChannel() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ch chan string) {
		i := 0
		defer wg.Done()
		for {
			if i == 5 {
				break
			}
			ch <- "abc"
			i++
		}
	}(ch)

	wg.Add(1)
	go func(ch chan string) {
		i := 0
		defer wg.Done()
		for {
			if i == 5 {
				break
			}
			fmt.Println(<-ch)
			i++
		}
	}(ch)

	wg.Wait()
}
