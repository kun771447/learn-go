package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁本质上是将并行的代码串行化，使用 lock 肯定会影响性能
// 即使是设计锁，那么也应该尽量的保证并行
// 我们有两组协程，其中一组负责写数据，另一组负责读数，web 系统大多场景读多写少
// 虽然有多个 roroutine，但是仔细分析我们发现，读协程之间应该并发，读和写之间应该串行(写的时候不能读)
func main() {
	var num int
	var rwlock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(2)
	// 写的 goroutine
	go func() {
		defer wg.Done()
		rwlock.Lock() // 加写锁，写锁会防止别的写锁和读锁获取
		defer rwlock.Unlock()
		num = 12
		time.Sleep(5 * time.Second)
		fmt.Println("get write lock", num)
	}()

	// 读的 goroutine
	go func() {
		defer wg.Done()
		for {
			rwlock.RLock() // 加读锁，读锁不会阻止别人的读
			time.Sleep(500 * time.Millisecond)
			fmt.Println("get read lock", num)
			rwlock.RUnlock()
		}
	}()

	wg.Wait()
}
