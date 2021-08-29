package main

import "fmt"

type Reader interface {
	ReadBook()
}
type Writer interface {
	WriteBook()
}

type Book struct {
}

func (b *Book) ReadBook() {
	fmt.Println("Book ReadBook")
}

func (b *Book) WriteBook() {
	fmt.Println("Book WriteBook")
}

func main() {
	var a string = "haha"
	// pair<type: string, value: "haha">
	fmt.Println(a)

	// pair<type: string, value: "haha">
	var allType interface{} = a
	fmt.Println(allType)

	value, ok := allType.(string)

	fmt.Println(value, ok)

	// pair<type: Book, value: book{} 地址>
	b := &Book{}

	// pair<type: Book, value: ''>
	var r Reader

	// pair<type: Book, value: book{} 地址>
	r = b

	r.ReadBook()

	var w Writer

	// pair<type: Book, value: book{} 地址>
	w = r.(Writer) // 断言可以成功？因为它们的 pair 是一致的

	w.WriteBook()
}
