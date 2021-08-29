package main

import (
	"fmt"
)

func main() {
	// 定义 channel

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 运行中")

		c <- 666 // 把 666 发送给 c，会阻塞
	}()

	number := <-c // 从 channel 中获取数据，会阻塞

	// 当 channel 中的内容未被完全取出或者完全写入，会造成阻塞

	fmt.Println("number", number)

	// for {
	// 	time.Sleep(time.Second)
	// }
}
