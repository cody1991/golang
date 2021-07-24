package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array)

	// for range
	for index, value := range array {
		fmt.Println("index", index, "value", value)
	}

}
