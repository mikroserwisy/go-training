package main

import (
	"math/rand"
	"sync"
	"fmt"
	"strconv"
	"time"
)

var store = make([]int, 0)
var group = sync.WaitGroup{}
var mutex = sync.RWMutex{}

func produce(name string) {
	for counter := 0; counter < 10; counter++ {
		mutex.Lock()
		store = append(store, rand.Int())
		fmt.Printf("P: %v: %v\n", name, store)
		mutex.Unlock()
	}
	group.Done()
}

func consume(name string)  {
	for counter := 0; counter < 10; counter++ {
		mutex.Lock()
		if (len(store) > 0) {
			store = store[:len(store) - 1]
			fmt.Printf("C: %v: %v\n", name, store)
		}
		mutex.Unlock()
	}
	group.Done()
}

func logger(name string)  {
	for {
		mutex.RLock()
		fmt.Printf("L: %v: %v\n", name, store)
		mutex.RUnlock()
		time.Sleep(1000)
	}
}

func main() {
	group.Add(20)
	go logger("L")
	for counter := 1; counter <= 10; counter++ {
		go produce("P" + strconv.Itoa(counter))
		go consume("C" + strconv.Itoa(counter))
	}
	group.Wait()
}