package main

import (
	"errors"
	"fmt"
)

// error, panic, recover
// 没有 try catch，捕获异常 这个概念
// 必须要处理这个 error, 返回一个值，告知是否成功, if err != nil

// panic 函数会导致程序退出，不推荐使用
// 使用场景：
// 比如我们要启动某些服务，需要准备某些依赖，比如日志文件，mysql 能连接通
// 发现一个不满足那就调用 panic 主动调用，终止程序

// recover 捕获 panic，防止程序终止
// defer 需要在 panic 前定义，另外 recover 只在 defer 中生效
// recover 处理异常后，逻辑并不会在 panic 那个点后继续执行
func test() (int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover", r)
		}
	}()

	panic("panic")
	return 0, errors.New("this is an error")
}
func main() {
	if _, err := test(); err != nil {
		fmt.Println(err)
	}
}
