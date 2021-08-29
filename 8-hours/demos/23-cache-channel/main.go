package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // 带有缓冲的 channel

	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子 go 结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子 go 正在运行,发送的元素是", i, "len(c) = ", len(c), "cap(c) = ", cap(c))
		}
	}()

	time.Sleep(time.Second)

	for i := 0; i < 4; i++ {
		num := <-c // 从 c 中取值
		fmt.Println("num = ", num)
	}

	fmt.Println("main go 结束")

	for {
		time.Sleep(time.Second)
	}
}
