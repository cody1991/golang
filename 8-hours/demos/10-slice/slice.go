package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	var slice2 []int

	if slice2 == nil {
		fmt.Println("slice2 is nil")
	}

	fmt.Printf("slice1 len: %d, slice = %v\n", len(slice1), slice1)

	// slice2[0] = 1 // 需要给它开辟空间

	slice2 = make([]int, 10)
	slice2[9] = 10

	fmt.Printf("slice2 len: %d, slice = %v\n", len(slice2), slice2)

	slice3 := make([]int, 10)

	fmt.Printf("slice3 len: %d, slice = %v\n", len(slice3), slice3)

}
