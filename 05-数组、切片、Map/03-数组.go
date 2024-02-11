package main

import "fmt"
func main() {
	// 数组定义 var name [count]int
	var arr []string;
	var arr2 [10]string;
	// []string 和 [3]string // 这是两种不同的类型

	fmt.Println("%T", arr);
	fmt.Println("%T", arr2);

	fmt.Println(arr);

	arr2[0] = "0";
	arr2[1] = "1";
	arr2[2] = "2";
	arr2[3] = "3";

	// arr2.append("4");

	fmt.Println("---", len(arr2));

	// for _, value := range arr2 {
	// 	fmt.Println(value);
	// }
	
	for i:= 0; i < 5; i ++ {
		if (i > 2) {
			continue;
		}
		fmt.Printf(arr2[i] + "-");
	}
}