package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o") // 返回一个 strings.Replacer，把每个 "#" 替换成 "o"
	fixed := replacer.Replace(broken)         // strings.Replacer 调用 方法，传递参数为要替换字符的字符串

	fmt.Println(fixed)
}
