package main

import "fmt"

var ga int = 3

// gd:=4 non-declaration statement outside function body 不能在函数外声明

func main() {
	var a int = 1

	fmt.Printf("a is %d, type of a is %T\n", a, a)

	var b = 2

	fmt.Printf("b is %d, type of b is %T\n", b, b)

	var c string

	fmt.Printf("c is %s, type of c is %T\n", c, c)

	d := 3.14

	fmt.Printf("d is %f, type of d is %T\n", d, d)

	var x, y float64 = 100.0, 200.0
	fmt.Printf("x is %f, type of x is %T\n", x, x)
	fmt.Printf("y is %f, type of y is %T\n", y, y)

	var xx, yy = 100.0, 200
	fmt.Printf("xx is %f, type of xx is %T\n", xx, xx)
	fmt.Printf("yy is %d, type of yy is %T\n", yy, yy)
}
