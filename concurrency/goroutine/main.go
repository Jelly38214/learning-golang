package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("当前可用的物理处理器: ", runtime.NumCPU())
	runtime.GOMAXPROCS(1) // 配置只有一个逻辑处理器给调度器使用
	// wg 用来等待程序完成
	// 计数加2, 表示要等待两个goroutine
	var wg sync.WaitGroup // 计数信号量
	wg.Add(2)

	fmt.Println("Start Goroutines")
	// 声明一个匿名函数,并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 声明一个匿名函数,并创建一个goroutine
	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait() // 如果WaitGroup的值大于0, Wait方法就会阻塞. 等待goroutine完成

	fmt.Println("\nTerminating Program")
}
