package main

import (
	"fmt"
)

func main() {
	i := 10

	if i > 10 {
		fmt.Println("i > 10")
	} else if i > 5 {
		fmt.Println("5 <= i <= 10")
	} else {
		fmt.Println("i < 5")
	}

	// 简单语句，j 只能在里面使用
	if j := 10; j > 20 {
		fmt.Println("j > 20")
	} else {
		fmt.Println("j <= 20")
	}

	// fmt.Println(j) undeclared name: j

	// 等同 if else
	switch i = 10; {
	case i > 10:
		fmt.Println("i > 10")
	case 5 <= i && i <= 10:
		fmt.Println("5 <= i <= 10")
	default:
		fmt.Println("i <5")
	}

	switch i := 1; i {
	case 1:
		fallthrough
	case 2:
		fmt.Println("1")
		fallthrough
	case 3:
		fmt.Println("2")
	default:
		fmt.Println("没匹配")
	}

	fmt.Println(i)

	switch 2 > 1 {
	case true:
		fmt.Println("2>1")
	case false:
		fmt.Println("2<=1")
	}

	sum := 0

	for i = 0; i < 100; i++ {
		sum += i
	}

	fmt.Println("sum", sum)

	sum2 := 0
	i = 0

	for i < 100 {
		sum2 += i
		i++
	}

	fmt.Println("sum2", sum2)

	sum3 := 0
	i = 0

	for {
		sum3 += i
		i++
		if i > 99 {
			break
		}
	}

	fmt.Println("sum3", sum3)

	sum4 := 0

	for i = 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		}
		sum4 += i
	}
	fmt.Println("sum4", sum4)

}
