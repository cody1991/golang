package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func(a int, b int) bool {

		fmt.Println("a,b", a, b)

		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")

			// go 协程退出
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")

		return true // 获取需要使用 channel
	}(10, 20)

	for {
		time.Sleep(time.Second)
	}
}
