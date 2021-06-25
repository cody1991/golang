package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	var myFloat float64
	var myBoolean bool

	fmt.Println(&myInt)
	fmt.Println(&myFloat)
	fmt.Println(&myBoolean)

	fmt.Println(reflect.TypeOf(&myInt))
	fmt.Println(reflect.TypeOf(&myFloat))
	fmt.Println(reflect.TypeOf(&myBoolean))

	var myIntPointer *int // 声明指向int指针的变量
	myIntPointer = &myInt // 给变量分配一个int指针
	fmt.Println(myIntPointer)

	myBoolean = true
	myBooleanPointer := &myBoolean
	fmt.Println(myBooleanPointer)
	fmt.Println(*myBooleanPointer)

	*myBooleanPointer = false
	fmt.Println(*myBooleanPointer)

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	pointPointer(&myBoolean)

	amount := 12
	double(&amount) // 传递指针
	fmt.Println(amount)

}

func createPointer() *float64 {
	myFloat := 98.5
	return &myFloat
}

func pointPointer(myBooleanPointer *bool) {
	fmt.Println(*myBooleanPointer)
}

// 接受指针参数
func double(number *int) {
	*number *= 2 // 更改指针处的值
}
