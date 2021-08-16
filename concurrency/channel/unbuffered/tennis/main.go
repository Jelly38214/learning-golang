package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int) // 创建一个无缓冲的通道

	wg.Add(2)
	go player("Nadal", court)
	go player("Djokovic", court)

	// 开始发球
	court <- 1 // 此时没有接收方, 程序阻塞

	wg.Wait()
}

// 网球比赛选手总是处于等待接球/将球打向对方的状态
// 使用两个goroutine来模拟网球比赛, 并使用无缓冲的通道来模拟球的来回
func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 等待球被击打过来, 接受值
		ball, ok := <-court // 锁住执行,等待有值发送或者通道关闭

		if !ok { // 通道关闭
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选随机数, 然后用这个数来判断我们是否丢球
		n := rand.Intn(100) // 产生[0, 100]的随机数

		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court) // 关闭通道, 表示没有球击打回去
			return
		}

		// 显示击球数, 并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// 击球回去
		court <- ball
	}
}
