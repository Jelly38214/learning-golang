package main

/**
 * 依据调度运行的无人值守的面向任务的程序,及其所使用的并发模式
 * 1. 程序可以在分配的时间内完成工作, 正常终止
 * 2. 程序没有及时完成工作, "自杀"
 * 3. 接收到操作系统发送的中断事件,程序立刻试图清理状态并停止工作
 */

// runner包用于展示如何使用通道来监视程序的执行时间
// 如果程序运行时间太长,也可以使用它来终止程序

// runner包管处理任务的运行和生命周期
import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

//Runner 在给定的超时时间内执行一组任务, 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	interrupt chan os.Signal   // 收发os.Signal接口类型的值,用来从主机操作系统接收中断事件
	complete  chan error       // 收发error接口类型值. 执行任务的goroutine用来发送任务已经的信号
	timeout   <-chan time.Time // 管理执行任务的时间,如果从这个通道接收到一个time.Time的值,这个程序就会试图清理状态并停止工作
	tasks     []func(int)      // 函数值的切片, 代表一个接一个顺序执行的函数. 会有一个与main函数分离的goroutine来执行这些函数
}

// ErrTimeout 会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New 返回一个新的准备使用的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d), // 运行时会在指定的duration时间到期之后,向这个通道发送一个time.Time的值
	}
}

// Add 将一个任务附加到Runner上, 这个任务是一个接收一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有任务, 并监视通道事件
func (r *Runner) Start() error {
	// 希望接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时发出信号
	case err := <-r.complete:
		return err
		// 当任务处理程序运行超时时发出信号
	case <-r.timeout:
		return ErrTimeout
	}
}

// run 执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		task(id)
	}

	return nil
}

// gotInterrupt 验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

// timeout 规定了必须多少秒内处理完成
const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")

	// 为本次执行分配超时时间
	r := New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
