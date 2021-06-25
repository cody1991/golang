package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	render := bufio.NewReader(os.Stdin)
	input, err := render.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)             // 去掉首尾的空格，类似js的trim
	grade, err2 := strconv.ParseFloat(input, 64) // 字符串转浮点数

	if err2 != nil {
		log.Fatal(err)
	}

	var status string

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println(status)
}
