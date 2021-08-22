package main

import "fmt"

type user struct {
	name  string
	email string
}

// notify 使用值接收者实现了一个方法
// 值接收者声明的方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail使用指针接收者实现了一个方法
// 指针接收者声明的方法,调用时,会创建对应的副本
func (u *user) changeEmail(email string) {
	u.email = email
}

// 总结: 值接收者使用值的副本来调用方法, 而指针接收者使用实际值来调用方法

func main() {
	// user类型的值可以用来调用使用值接收者声明的方法
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// 指向user类型值的指针也可以用来调用使用值接收者声明的方法
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// user类型的值可以用来调用使用指针接收者声明的方法
	bill.changeEmail("bill@newdomain.com") // Go语言在背后对值做了调整,使得满足函数的接收者,进行调用(&bill.changeEmail("xx"))
	bill.notify()

}
