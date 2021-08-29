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
		"bg": "北京",
	}
	fmt.Println(myMap3)

	printMap(myMap3)

	delete(myMap3, "sz")

	printMap(myMap3)

}

func printMap(myMap map[string]string) {
	// myMap 是引用传递
	for index, value := range myMap {
		fmt.Println(index, value)
	}

	myMap["gz"] = "广州"
}
