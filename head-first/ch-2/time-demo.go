package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now() // 返回当前日期时间的 time.Time 值
	var year int = now.Year()      // 返回年份

	fmt.Println(year)
}
