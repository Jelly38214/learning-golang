package main

import "fmt"

// Person 声明一个结构体
type Person struct {
	Name string
	age  int
}

// Books结构体
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// 初始化结构体
	// 1. 按照顺序提供初始化值
	p1 := Person{"Tom", 25}
	// 2. 通过field:value的方式初始化,这样可以任意顺序
	p2 := Person{age: 24, Name: "Tom"}
	// 3. new方式,未设置初始值的,会赋予类型的零值
	p3 := new(Person)

	fmt.Println(p1, p2, p3)

	// 通过点.操作符访问,修改结构体
	var (
		Book1, Book2 Books
	)

	/* book1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "Jelly"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 123

	/* book2 描述 */
	Book2.title = "Python 语言"
	Book2.author = "Jelly"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 456

	/* 打印Book1信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印Book2信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
}
