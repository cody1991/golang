package main

import "fmt"

// 方法必须有一个接收者，接收者是一个类型
// 方法和类型绑定在一起，称为类型的方法
type Age uint

// age Age 是接收者
// age 是变量名，Age是接收者类型
func (age Age) String() {
	fmt.Println("the age is ", age)
}

// 指针类型可以修改值
func (age *Age) Modify(value Age) {
	*age = Age(value)
}

func main() {
	age := Age(1)
	age.String()

	sm := Age.String

	// 第一个参数必须是接收者
	sm(age)

	age.Modify(20)
	age.String()
}
