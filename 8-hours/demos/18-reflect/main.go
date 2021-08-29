package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id   int
	name string
	age  int
}

func (u *User) Call() {
	// 换成 User 才能打印出来，=。=啥原因
	fmt.Printf("%v\n", u)
}

func reflectNum(arg interface{}) {
	fmt.Println("type of arg", reflect.TypeOf(arg))
	fmt.Println("value of arg", reflect.ValueOf(arg))
}

func reflectMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)

	fmt.Println("reflectMethod", inputType.Name(), inputValue)

	// input 类型的 field 数量： NumField()
	for i := 0; i < inputType.NumField(); i++ {

		// 取得第 i 个 field
		field := inputType.Field(i)
		// 取得第 i 个 field 的值
		value := inputValue.Field(i)

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v", m.Name, m.Type)
	}
}

func main() {
	num := 1.2

	reflectNum(num)

	user := User{
		id:   1,
		name: "haha",
		age:  30,
	}

	user.Call()

	reflectMethod(user)
}
