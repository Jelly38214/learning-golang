package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// 传统的同步goroutine的机制, 就是对共享资源加锁
// atomic 和 sync包里面的函数提供了很好的解决方案

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final Counter: ", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 安全地对counter加1
		atomic.AddInt64(&counter, 1) // 强制同一时刻只能有一个goroutine运行并完成这个加法操作

		// 当前goroutine从线程退出,并放回到队列
		runtime.Gosched()
	}
}
