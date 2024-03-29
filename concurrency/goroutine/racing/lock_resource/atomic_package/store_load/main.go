package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

// atomic包里的Store和Load类函数来提供对数值类型的安全访问: 读与写
func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	// 给定goroutine执行的时间
	time.Sleep(1 * time.Second)

	// 该停止工作了, 安全地设置shutdown标志
	fmt.Println("Shutting Now")
	atomic.StoreInt64(&shutdown, 1)

	// 等待goroutine结束
	wg.Wait()
}

// doWork 用来模拟执行工作的goroutine
// 检测之前的shutdown标志来决定是否提前终止
func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// 要停止工作了吗?
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Donw\n", name)
			break
		}
	}
}
