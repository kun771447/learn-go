package test

import "fmt"

func (t Test) GetName() string {
	fmt.Println(t.Name)
	return t.Name
}
