package main

import "fmt"

// interface {} 通用万能接口，int struct float64 string 等等都实现了它

func myFunc(arg interface{}) {
	fmt.Println("\narg", arg)
	fmt.Printf("type of arg is %T\n", arg)

	// interface{} 如何知道底层的类型？
	// 使用断言机制

	value, ok := arg.(string)

	if !ok {
		fmt.Println("arg is not a string")
	} else {
		fmt.Println(value, "is a string")
	}
}

type Book struct {
	name string
}

func main() {
	book := Book{
		name: "haha",
	}
	myFunc(book)
	myFunc(1)
	myFunc("abcc")
}
