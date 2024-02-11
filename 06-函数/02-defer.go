package main

import (
	"fmt"
)

func main() {
	// 链接数据库，打开文件，开始锁，
	// 最后一定要关闭
	// var mu sync.Mutex

	// mu.Lock()
	// defer mu.Unlock()

	//  defer 后面的代码是会放在函数 return 之前执行
	// 多个 defe 会形成栈，后定义的 defer 先执行
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")

	fmt.Println("main")

	return
}
