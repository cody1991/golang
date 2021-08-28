package main

import (
	"fmt"
	"golang/8-hours/demos/6-package/lib1"
	"golang/8-hours/demos/6-package/lib2"
)

func main() {
	fmt.Println("main")

	lib1.Api()
	lib2.Api()
}
