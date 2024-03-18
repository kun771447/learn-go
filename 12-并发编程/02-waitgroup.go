package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	// 监控多少个 goroutine 执行完毕
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println(i)
		}(i)
	}

	// 等待所有 goroutine 执行完毕
	wg.Wait()
	fmt.Println("main grountine")
}
