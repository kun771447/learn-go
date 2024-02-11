package main

import (
	"errors"
	"fmt"
)

// go 语言中全部是值传递, 会对参数进行拷贝
func add(a, b int) (int, error) {
	return a + b, nil
}

func call(op string) func() {

	return func() {
		fmt.Println("call" + op)
	}
}

// 错误处理的理念，一个函数可能出错
// 没有 try catch
// 要求返回 error
// 开发者必须处理 error 代码里有大量的 if err != nil

func A() (int, error) {
	return 0, errors.New("error")
}
func main() {
	// go 函数支持普通函数、匿名函数、闭包
	// 函数一等公民，可以做变量、参数、返回值
	fmt.Println(add(1, 3))

	if _, err := A(); err != nil {
		fmt.Println(err)
	}
}
