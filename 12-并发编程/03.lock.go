package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 锁-资源竞争

var total int
var total2 int32
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		atomic.AddInt32(&total2, 1)
		lock.Lock()
		total++
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		atomic.AddInt32(&total2, -1)
		lock.Lock()
		total--
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
	fmt.Println("total2", total2)
}
