package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			// panic: send on closed channel
			// close(c) // 关闭 channel
		}
		// 需要添加，all goroutines are asleep - deadlock!
		close(c) // 关闭 channel
	}()

	// for {
	// 	// ok true 表示channel没有关闭，false 表示已经关闭且无数据
	// 	if data, ok := <-c; ok {
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }
	// 另外一种简写的方式，使用range不断迭代操作channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("finish")
}
