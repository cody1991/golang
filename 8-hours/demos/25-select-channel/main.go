package main

import "fmt"

func fib(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			// 如果 c 可写，执行
			x, y = y, x+y
		case <-quit:
			// 如果 quit 可读，执行
			fmt.Println("quit")
			return
		}
	}
}

func main() {

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // 从 c 中读数据
		}

		quit <- 0
	}()

	fib(c, quit)

}
