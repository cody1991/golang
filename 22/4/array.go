package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array)

	// for range
	for index, value := range array {
		fmt.Println("index", index, "value", value)
	}

	array2 := [3][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(array2)
}
