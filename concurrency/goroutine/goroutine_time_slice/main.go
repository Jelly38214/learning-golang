package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 基于调度器内部算法, 一个正运行的goroutine在工作结束前, 可以停止并重新调度
// 调度器这样做的目的是防止某个goroutine长时间占用逻辑处理器.当goroutine占用时间过长时,调度器会停止当前的goroutine,并给其他可运行的goroutine运行的机会

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create Goroutine")

	// 由于printPrime耗时较长且只有一个逻辑处理器
	// 因此A/B两个Goroutine会交替执行
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("Terminating Program")

}

// printPrime 显示5000以内的素数值
func printPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}

		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)
}
