package main

import "fmt"

func deferFunc() int {
	fmt.Println("deferFunc")
	return 0
}

func returnFunc() int {
	fmt.Println("returnFunc")
	return 0
}

func returnAndDeferFunc() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	// defer 函数体结束之前触发 defer

	defer fmt.Println("end1")
	defer fmt.Println("end2")

	fmt.Println("hello1")
	fmt.Println("hello2")

	returnAndDeferFunc()
}
