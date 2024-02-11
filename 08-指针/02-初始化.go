package main

import "fmt"

type Person struct {
	name string
}

func changeName(p *Person) {
	p.name = "jack"
}
func main() {
	// 指针需要初始化

	// var b *int	// nil
	// var p *Person // nil

	// 第一种初始化方式
	p1 := &Person{}

	// 第二种初始化方式
	var emptyPerson Person
	p2 := &emptyPerson

	// 第三种初始化方式
	var p3 *Person = new(Person)

	// map、channe、slice 初始化推荐使用 make 方法
	// 指针推荐使用 new 函数，指针要初始化否则会出现 nil

	fmt.Println(p1, p2, p3)
}
