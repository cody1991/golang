package main

import "fmt"

// iota 只能配合 const() 使用

const (
	SHENZHEN = 0
	SHANGHAI = 1
	BEIJING  = 2
)

const (
	// 匹配都是每行两个变量
	a, b = iota + 1, iota + 2 // 1 2
	c, d                      // 2 3，沿用上面的规则， iota = 1
	e, f                      // 3 4，沿用上面的规则， iota = 2

	g, h = iota * 2, iota * 3 // 6 9 iota = 3，修改了公式
	i, j                      // 8 12 使用修改后的公式
)

const (
	// 每行的 iota 都会加 1，默认是 0
	SZ = 10 * iota // iota = 0
	SH             // iota = 1
	BG             // iota = 2
)

func main() {
	const length int = 100

	fmt.Println("length is ", length)
	fmt.Println("shenzhen is ", SHENZHEN)

	fmt.Println("SZ is ", SZ)
	fmt.Println("SH is ", SH)
	fmt.Println("BG is ", BG)

	fmt.Println("a,b,c,d,e,f,g,h,i,j", a, b, c, d, e, f, g, h, i, j) // 1 2 2 3 3 4 6 9 8 12

}
