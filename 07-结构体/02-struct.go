package main

import "fmt"

// 类型集合
var persons [][]string

type Person struct {
	name    string
	age     int
	address string
	height  float32
}

// 结构体定义方法
// 接收器有两种形态，
// 指针，修改 p 的值, 或者数据很大
func (p *Person) print() {
	p.age = 19
	fmt.Println("Person:", p.name, p.age, p.address, p.height)
}

// 结构体嵌套
type Student struct {
	// 具名嵌套
	p Person

	// 匿名嵌套
	Person

	p2 struct {
		name    string
		age     int
		address string
		height  float32
	}
}

func main() {
	// persons := append(persons, []string{"John", "Doe"})
	// fmt.Println(persons)
	p1 := Person{
		name:    "John",
		age:     20,
		address: "Shanghai",
		height:  170.5,
	}

	fmt.Println(p1)

	p1.print()

	fmt.Println(p1)

	// 匿名结构体，匿名函数
	address := struct {
		name string
		city string
	}{
		name: "北京",
	}

	fmt.Println(address.name)
	fmt.Println(address.city == "")
}
