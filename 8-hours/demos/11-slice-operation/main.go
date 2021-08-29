package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5) //第三个参数是容量

	fmt.Printf("len %d, cap %d, value %v, address %p\n", len(numbers), cap(numbers), numbers, numbers)

	numbers = append(numbers, 5, 10, 5)

	fmt.Printf("len %d, cap %d, value %v, address %p\n", len(numbers), cap(numbers), numbers, numbers)

	// 截取

	s1 := numbers[1:5] // 左闭右开
	s1[0] = 1000
	fmt.Println(s1)
	fmt.Println(numbers[:])
	fmt.Println(numbers[:3])
	fmt.Println(numbers[3:])

	s2 := make([]int, 2)
	copy(s2, s1)

	fmt.Println(s2)
}
