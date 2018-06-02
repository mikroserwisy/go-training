package main

import (
	"sync"
	"runtime"
	"sync/atomic"
)

var counter int64

var group sync.WaitGroup
//var mutex = sync.Mutex{}

func init()  {
	println("Init...")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func add()  {
	for count := 0; count < 10000; count++ {
		//mutex.Lock()
		//counter++
		//mutex.Unlock()
		atomic.AddInt64(&counter, 1)
	}
	group.Done()
}

func main() {
	group.Add(2)
	go add()
	go add()
	group.Wait()
	println(counter)
}



