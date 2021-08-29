package main

import (
	"fmt"
	"reflect"
)

type Resume struct {
	Name string `json:"name" doc:"我的名字"`
	Sex  string `json:"sex" doc:"我的性别"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		jsonInfo := t.Field(i).Tag.Get("json")
		docInfo := t.Field(i).Tag.Get("doc")
		fmt.Println(jsonInfo)
		fmt.Println(docInfo)
	}
}

func main() {
	var re Resume

	findTag(&re)
}
