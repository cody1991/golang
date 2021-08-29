package main

import "fmt"

// 声明一种新的数据类型，是 int 别名
type myInt int

// 定义结构体
type Book struct {
	title  string
	author string
}

func (book *Book) GetTitle() string {
	// 传递的是副本，使用 *Book
	return book.title
}

func (book *Book) GetAuthor() string {
	// 传递的是副本，使用 *Book
	return book.author
}

func (book *Book) SetTitle(title string) {
	// 传递的是副本，使用 *Book
	book.title = title
}

func changeBook(book *Book) {
	// 需要按引用调用
	book.title = "new golang"
}

func main() {
	var a myInt = 10

	fmt.Printf("typeof a: %T\n", a)

	book := Book{
		author: "zhangsan",
		title:  "golang",
	}

	book.author = "zhangsan"
	book.title = "golang"

	fmt.Printf("%v\n", book)

	changeBook(&book)

	fmt.Printf("%v\n", book)

	fmt.Printf("%v\n", book.GetTitle())
	fmt.Printf("%v\n", book.GetAuthor())

	book.SetTitle("new golang title")
	fmt.Printf("%v\n", book.GetTitle())
}
