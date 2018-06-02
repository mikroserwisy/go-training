package main

import (
	"log"
	"math/rand"
	"sync"
	"runtime"
)

var group sync.WaitGroup

func init()  {
	println("Init...")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func task(name string)  {
	for count := 0; count < 10000; count++ {
		log.Printf("%v: %v\n", name, rand.Int())
	}
	group.Done()
}

func main() {
	group.Add(2)
	go task("T1")
	go task("T2")
	group.Wait()
}
