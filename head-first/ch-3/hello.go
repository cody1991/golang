package main

import "fmt"

var metersPerLiter float64

func main() {

	metersPerLiter = 10.0

	fmt.Printf("one-third %0.2f\n", 1.0/3)

	// 返回格式化后的字符串
	oneThird := fmt.Sprintf("one-third %0.2f", 1.0/3)
	fmt.Println(oneThird)

	// %0.2f 格式化动词

	fmt.Printf("float: %f\n", 3.1415)
	fmt.Printf("float: %10.2f\n", 3.1415)
	fmt.Printf("float: %10f\n", 3.1415) // 宽度
	fmt.Printf("integer: %d\n", 15)
	fmt.Printf("string: %s\n", "hello")
	fmt.Printf("integer: %v %v %v\n", false, 15, 3.1415)

	// %#v可以显示一些值，如果不使用%#v的话，这些值可能会在输出中被隐藏，例如，在这段代码中，%#v显示了一个空字符串、一个制表符和一个换行符，所有这些在用%v打印时都是不可见的
	fmt.Printf("integer: %#v %#v %#v\n", "", "\t", "\n")
	fmt.Printf("integer: %v %v %v\n", "", "\t", "\n")
	fmt.Printf("integer: %T %T %T\n", false, 15, 3.1415)
	fmt.Printf("float: %%\n")

	sayHello()

	repeatLine("codytang", 3)

	// 打印格式化的字符串

	fmt.Printf("area is %10.2f\n", printNeeded(1.0, 2.0))
	fmt.Printf("area is %10.2f\n", printNeeded(1.12312, 2.231))

}

func printNeeded(width float64, height float64) float64 {
	area := width * height
	return area / metersPerLiter
}

func sayHello() {
	fmt.Println("Hello")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}
