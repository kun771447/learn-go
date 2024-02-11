package main

import (
	"container/list";
	"fmt";
)

func main() {
	var myList list.List;
	// PushBack(); 		// 尾部
	// PuhstFront(); 	// 头部
	// InsertBefore(value， index); 	// 插入

	myList.PushBack(1);
	myList.PushBack(2);
	e3 := myList.PushBack(3);


	myList.InsertBefore(4, e3);

	fmt.Println(myList);

	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value);
	}
}