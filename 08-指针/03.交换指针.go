package main

import "fmt"

type Person struct {
	name string
}

func swap(a, b *int) {
	c := *b
	*b = *a
	*a = c

	fmt.Println(*a, b)
}

func main() {
	a, b := 1, 2

	swap(&a, &b)

	fmt.Println(a, b)
}
