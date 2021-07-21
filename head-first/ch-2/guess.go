// 猜数字
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix() // 时间戳
	// fmt.Println(seconds)

	/***
	为了得到不同的随机数，我们需要向rand.Seed函数传递一个值。
	这将“播种”随机数生成器，也就是说，给它一个值，它将用于生成其他随机值。
	但如果我们一直给它相同的种子值，它就会一直给我们相同的随机值，我们会回到开始的地方。
	我们之前看到过。time.Now函数会提供一个表示当前日期和时间的Time值。每次运行程序时，我们都可以使用它来获得不同的种子值。
	**/
	rand.Seed(seconds)

	// rand.Intn 它将返回一个介于0和你提供的数字之间的随机整数
	// 如果传递一个100的参数，我们将得到一个0~99范围内的随机数
	target := rand.Intn(100) + 1
	// fmt.Println(target)

	render := bufio.NewReader(os.Stdin)

	limit := 10
	success := false
	for guesses := 0; guesses < limit; guesses++ {
		fmt.Println("you have", limit-guesses, "guesses left")

		fmt.Print("Make a guess: ")
		input, err := render.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		grade, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(grade)

		if grade < target {
			fmt.Println("Your guess was LOW")
		} else if grade > target {
			fmt.Println("Your guess was HIGH")
		} else {
			fmt.Println("Good Job!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("target is", target)
	}
}
