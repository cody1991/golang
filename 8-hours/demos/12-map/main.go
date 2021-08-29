package main

import "fmt"

func main() {
	// key 是 字符串， value 也是字符串
	var myMap map[string]string

	fmt.Println(myMap)

	myMap = make(map[string]string, 10)

	fmt.Println(myMap)

	myMap["sz"] = "深圳"

	fmt.Println(myMap)

	myMap2 := make(map[string]string, 10)
	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"sz": "深圳",
	}
	fmt.Println(myMap3)

}
