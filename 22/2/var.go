package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// var 变量名 类型 = 表达式

	var i int = 10
	fmt.Println(i)

	var j = 10
	fmt.Println(j)

	var (
		k int = 20
		l int = 30
	)

	fmt.Println(k, l)

	var (
		a = 10
		b = 20
		c = 30
	)

	fmt.Println(a, b, c)

	var f32 float32 = 2.2
	var f64 float64 = 2.2345678
	fmt.Println("f32 = ", f32, "f64 = ", f64)

	var bt bool = true
	var bf bool = false

	fmt.Println("bt", bt, "bf", bf)

	var s1 string = "hello"
	var s2 string = "world"

	fmt.Println("s1", s1, "s2", s2)
	fmt.Println("s1+s2=", s1+s2)

	var zi int
	var zf float64
	var zb bool
	var zs string

	fmt.Println(zi, zf, zb, zs)

	si := 20
	ss := "123"

	fmt.Println(si, ss)

	pi := &si

	fmt.Println(pi, *pi)

	const name = "codytang"

	const (
		one = iota + 3
		two
		three
		four
	)

	fmt.Println(one, two, three, four)

	// int 类型转成 str 类型
	i2s := strconv.Itoa(si)

	fmt.Println(i2s)

	// str 类型转为 int 类型
	s2i, err := strconv.Atoi(ss)

	fmt.Println(s2i, err)

	// 数字 浮点互转
	i2f := float64(si)
	f2i := int(i2f)

	fmt.Println(i2f, f2i)

	fmt.Println(strings.HasPrefix(name, "cody"))
	fmt.Println(strings.Index(name, "cody"))
	fmt.Println(strings.ToUpper(name))

}
