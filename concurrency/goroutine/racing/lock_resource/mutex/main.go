package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 同步访问共享资源的另一种方式是使用互斥锁
// 互斥锁用于在代码上创建一个临界区,保证同一个时间只有一个goroutine可以执行这个临界区代码

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex // mutex 用来定义一段代码临界区
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Final Counter: %d", counter)
}

// incCounter 使用互斥锁来同步保证安全访问
func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock() // 释放锁,允许其他正在等待的goroutine进入临界区
	}
}
