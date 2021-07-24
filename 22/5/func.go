package main

import (
	"errors"
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

// 可以省略相同变量类型
func sum2(a, b int) int {
	return a + b
}

func sum3(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能小于0")
	}
	return a + b, nil
}

func sum4(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能小于0")
	}
	sum = a + b
	err = nil
	return sum, err
}

func sum5(params ...int) int {
	sum := 0

	for _, value := range params {
		sum += value
	}

	return sum
}

// 必包：返回 函数，它的返回值是 int
func colsure() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	result := sum(1, 2)
	fmt.Println(result)

	result2, err := sum3(-1, -2)

	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("result2", result2)
	}

	fmt.Println(sum4(1, 2))
	fmt.Println(sum4(1, -2))

	fmt.Println(sum5(1, 2, 3))

	sum6 := func(a, b int) int {
		return a + b
	}

	fmt.Println(sum6(1, 2))

	c := colsure()

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

}
