package main

import (
	"fmt"
	"sync"
	"time"
)

// each fork has a mutex, so each individual fork can be locked or unlocked
type Fork struct {
	sync.Mutex
}

// generate 5 forks
func getForks() [5]*Fork {
	arr := [5]*Fork{}
	for i := 0; i < 5; i++ {
		arr[i] = &Fork{sync.Mutex{}}
	}
	return arr
}

var forks = getForks()

// each philosopher only has a id associated with them
var philosophers = [5]int{0, 1, 2, 3, 4}

// this func allows only 2 philosophers to eat concurrently
// for this philosopher takes lock on left and right fork first before eating
// once he finished eating, he releases the lock
func eat(philosopherId int, wg *sync.WaitGroup) {
	defer wg.Done()
	leftFork := (philosopherId + 4) % 5
	rightFork := (philosopherId + 1) % 5

	forks[leftFork].Lock()
	forks[rightFork].Lock()

	fmt.Println("Philosopher ", philosopherId, " is eating")

	// assuming each philoshoper will take 2 sec in eating
	time.Sleep(2 * time.Second)

	forks[leftFork].Unlock()
	forks[rightFork].Unlock()

	fmt.Println("Philosopher ", philosopherId, "finished eating")
}

func main() {
	wg := &sync.WaitGroup{}

	for j := 0; j < 5; j++ {
		wg.Add(1)
		go eat(philosophers[j], wg)
	}

	wg.Wait()
}
