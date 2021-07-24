package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}

	slice := array[2:4]

	fmt.Println(slice)

	for index, value := range slice {
		fmt.Println("index", index, "value", value)
	}

	slice[0] = "modify"

	fmt.Println(array, slice)

	// 4 是长度，8是容量
	slice1 := make([]string, 4, 8)
	fmt.Println(slice1)

	fmt.Println(len(slice), cap(slice))

	slice2 := append(slice, "f", "g")

	fmt.Println(slice2)
}
