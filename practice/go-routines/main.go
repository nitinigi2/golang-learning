package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = make(map[int]Book)
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		id := rnd.Intn(10) + 1
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
			m.Lock()
			b, ok := cache[id]
			m.Unlock()
			if ok {
				fmt.Println("From cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)

		go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
			b, ok := queryDb(id, m)
			if ok {
				fmt.Println("From database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryDb(id int, m *sync.Mutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range Books {
		if b.ID == id {
			m.Lock()
			cache[b.ID] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
