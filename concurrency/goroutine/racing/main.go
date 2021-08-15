package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int // 所有goroutine都要增加其值的变量
	wg      sync.WaitGroup
)

// 竞态导致并发程序变得复杂,难以调试
func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter: ", counter) // ?
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		// 当前goroutine从线程退出,并放回到队列
		runtime.Gosched()

		value++
		counter = value
	}
}
