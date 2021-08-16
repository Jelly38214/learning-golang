package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutine = 4  // 要使用的goroutine的数量
	taskLoad        = 10 // 要处理的工作的数量
)

var wg sync.WaitGroup

// 展示如何使用有缓冲的通道和固定数目的goroutine来处理一堆工作

// init初始包,Go语言运行时会在其他代码执行前优先执行这个函数
func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutine)

	// 创建worker
	for gr := 1; gr <= numberGoroutine; gr++ {
		go worker(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 当所有工作都处理完成时关闭通道,以便所有goroutine退出
	close(tasks) // 关闭通道, 但依旧可以从其中接收数据(只要有),只是不能在往其发送

	wg.Wait()
}

// worker作为goroutine启动来处理从有缓冲的通道闯入
func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks // 取出task

		if !ok {
			// 这意味着通道已经空, 并且已被关闭
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// 显示开始工作
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// 随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示完成工作
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
