package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (h *Human) Eat() {
	fmt.Println(h.name, "Human Eat")
}

func (h *Human) Walk() {
	fmt.Println(h.name, "Human Walk")
}

type Super struct {
	Human
	level int
}

func (s *Super) Eat() {
	fmt.Println(s.name, "Super Eat")
}

func (s *Super) Fly() {
	fmt.Println(s.name, "Super Fly")
}

func main() {
	h := Human{
		name: "张三",
		sex:  "male",
	}

	h.Walk()
	h.Eat()

	s := Super{
		Human{
			name: "李四",
			sex:  "male",
		},
		88,
	}

	s.Fly()
	s.Eat()
	s.Walk()
}
