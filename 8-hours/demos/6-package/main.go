package main

import (
	"fmt"
	myLib1 "golang/8-hours/demos/6-package/lib1" // 别名
	_ "golang/8-hours/demos/6-package/lib2"      // 匿名 _ ，导入不使用，会调用包的 init 方法
	// . "golang/8-hours/demos/6-package/lib2"  // 直接导入整个包，可以使用里面的方法。可能会有重复的冲突
)

func main() {
	fmt.Println("main")

	myLib1.Api()
	// lib2.Api()
	// Api()
}
