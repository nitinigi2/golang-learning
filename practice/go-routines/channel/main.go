package main

import (
	"fmt"
	"sync"
)

func main() {
	// wg := &sync.WaitGroup{}
	// ch := make(chan int)

	// wg.Add(2)

	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	fmt.Println(<-ch)
	// 	wg.Done()
	// }(ch, wg)

	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	ch <- 10
	// 	wg.Done()
	// }(ch, wg)

	// close(ch)

	// wg.Wait()

	SelectExample()
}

func BufferedChannel() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 10
		ch <- 20
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func channelType() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)

	// recieve only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	// send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 10
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
