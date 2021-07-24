package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	nameAgeMap := make(map[string]int)

	nameAgeMap["cody"] = 20
	nameAgeMap["tang"] = 20

	fmt.Println(nameAgeMap)

	nameAgeMap2 := map[string]int{"hello": 30}

	nameAgeMap2["cody"] = 20
	nameAgeMap2["tang"] = 20

	fmt.Println(nameAgeMap2)

	age, ok := nameAgeMap2["cody"]

	fmt.Println(age, ok)

	delete(nameAgeMap2, "cody")

	fmt.Println(nameAgeMap2)

	for key, value := range nameAgeMap2 {
		fmt.Println(key, value)
	}

	fmt.Println(len(nameAgeMap2))

	s := "codytangahhh 唐泽雄"
	bs := []byte(s)

	fmt.Println(bs)
	fmt.Println(len(bs))

	// 汉字不当做3个
	fmt.Println(utf8.RuneCountInString(s))

	for i, value := range s {
		fmt.Println(i, value)
	}
}
