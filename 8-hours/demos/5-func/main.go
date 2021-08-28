package main

import "fmt"

func sum(a int, b int) int {
	fmt.Println("a", a)
	fmt.Println("b", b)

	c := a + b

	return c
}

// 匿名返回值
func calc(a int, b int) (int, int) {
	fmt.Println("a", a)
	fmt.Println("b", b)
	return a + b, a - b
}

// 行参名返回值
func calc2(a int, b int) (sum int, sub int) {
	fmt.Println("a", a)
	fmt.Println("b", b)
	sum = a + b
	sub = a - b
	return // 可以不写返回的值，按照行参名的顺序
}

// 行参名返回值，省略一个 int
func calc3(a int, b int) (sum, sub int) {
	fmt.Println("a", a)
	fmt.Println("b", b)
	sum = a + b
	sub = a - b
	return // 可以不写返回的值，按照行参名的顺序
}

func main() {
	fmt.Println("sum", sum(1, 2))
	calcSum, calcSub := calc(1, 2)
	calcSum2, calcSub2 := calc2(1, 2)
	calcSum3, calcSub3 := calc3(1, 2)

	fmt.Println("calcSum", calcSum, "calcSub", calcSub)
	fmt.Println("calcSum2", calcSum2, "calcSub2", calcSub2)
	fmt.Println("calcSum3", calcSum3, "calcSub3", calcSub3)
}
