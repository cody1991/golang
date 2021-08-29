package main

import "fmt"

// 本质是指针
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
	name  string
}

func (c *Cat) Sleep() {
	fmt.Println(c.name, "Cat Sleep")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "cat"
}

type Dog struct {
	color string
	name  string
}

func (d *Dog) GetColor() string {
	return d.color
}

func (d *Dog) GetType() string {
	return "dog"
}

func (d *Dog) Sleep() {
	fmt.Println(d.name, "Dog Sleep")
}

func showAnimal(animal Animal) {
	animal.Sleep()

	fmt.Println(animal.GetType())
	fmt.Println(animal.GetColor())
}

func main() {

	// 一个父类 （接口）
	// 多个子类（实现了父类的所有接口）
	// 父类类型的变量（指针）指向(引用) 子类的具体数据变量
	var animal Animal

	animal = &Cat{
		name:  "mimi",
		color: "black",
	}

	fmt.Println(animal)

	animal.Sleep()

	cat := Cat{
		name:  "haha",
		color: "black",
	}

	dog := Dog{
		name:  "hahaha",
		color: "while",
	}

	showAnimal(&cat)
	showAnimal(&dog)

}
