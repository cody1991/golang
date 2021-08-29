package main

import "fmt"

func changeValue(b *int) {
	*b = 20
}

func main() {
	var a int = 10

	changeValue(&a)

	fmt.Println("a", a)

	b, c := 10, 20

	swap(&b, &c)

	fmt.Println("b,c", b, c)

	var p *int = &a

	fmt.Println("p,&a", p, &a)

	var pp **int = &p

	fmt.Println("pp,&p", pp, &p)
}

func swap(pa *int, pb *int) {
	tmp := *pa // 取得 A 的值
	*pa = *pb  // 取得 B 的值，赋值给 A
	*pb = tmp  // A 的值 赋值给 B
}
