package main

import "fmt"

func main() {
	var a [10]int

	fmt.Println("a", a)

	a[5] = 20

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	b := [4]int{1, 2, 3, 4}
	fmt.Println(b)

	// 关键字 range
	for index, value := range b {
		fmt.Println(index, value)
	}

	fmt.Printf("b type is %T\n", b)

	printArray(b)

	var c = []int{1, 2, 3, 4} // 动态数组，切片slice
	printArray2(c)            // 切片是使用引用传递
	printArray2(c)
}

func printArray(arr [4]int) {
	// 数组也是值传递
	for index, value := range arr {
		fmt.Println(index, value)
	}
}

func printArray2(arr []int) {
	// 数组也是值传递
	for index, value := range arr {
		fmt.Println("printArray2", index, value)
	}

	arr[0] = 20
}
