package main

import (
	"fmt"
	"sync"
	"time"
)

type Fork struct {
	available bool
	m         sync.Mutex
}

func getForks() [5]*Fork {
	arr := [5]*Fork{}
	for i := 0; i < 5; i++ {
		arr[i] = &Fork{
			available: true,
		}
	}
	return arr
}

var forks = getForks()

var philosophers = [5]int{0, 1, 2, 3, 4}

func main() {
	wg := &sync.WaitGroup{}

	for j := 0; j < 5; j++ {
		wg.Add(1)
		go eat(philosophers[j], wg)
	}

	wg.Wait()
}

func eat(philosopherId int, wg *sync.WaitGroup) {
	defer wg.Done()
	left := (philosopherId + 4) % 5
	right := (philosopherId + 1) % 5

	forks[left].m.Lock()
	forks[right].m.Lock()

	forks[left].available = false
	forks[right].available = false

	fmt.Println("Philosopher ", philosopherId, " is eating")
	time.Sleep(time.Millisecond * time.Duration(philosopherId+900))
	fmt.Println("Philosopher ", philosopherId, "finished eating")

	forks[left].available = true
	forks[right].available = true

	forks[left].m.Unlock()
	forks[right].m.Unlock()
}
