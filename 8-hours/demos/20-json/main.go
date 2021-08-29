package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "张柏芝"}}

	fmt.Println(movie)

	// 编码过程 struct => json
	jsonStr, err := json.Marshal(movie)

	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	fmt.Printf("json str %s\n", jsonStr)

	// 解码过程 json => struct
	var newMove Movie
	err = json.Unmarshal(jsonStr, &newMove)

	if err != nil {
		fmt.Println("json Unmarshal error", err)
		return
	}

	fmt.Printf("%v\n", newMove)
}
