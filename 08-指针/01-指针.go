package main

import "fmt"

type Person struct {
	name string
}

func changeName(p *Person) {
	p.name = "jack"
}
func main() {
	// 指针

	p := Person{
		name: "tom",
	}

	changeName(&p)

	fmt.Println(p)

	po := &Person{
		name: "tom",
	}
	(*po).name = "*tom"

	po.name = "tom2"

	fmt.Println(po)
}
