package main

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// pool 包用于展示如何使用有缓冲的通道实现资源池,来管理可以在任意数量的goroutine之间共享及独立使用的资源
// goroutine需要从池里得到这些资源中的一个,从池里申请,使用完后归还回资源池
// 标准库的实现: sync.Pool

// Pool 管理一组可以安全地在多个goroutine间共享的资源
// 被管理的资源必须实现io.Closer接口
type Pool struct {
	m         sync.Mutex                // 互斥锁保证多个goroutine访问资源池时,资源池的值时安全的
	resources chan io.Closer            // 只要某类资源实现了io.Closer接口,就可以用这个资源池来管理
	factory   func() (io.Closer, error) // 当池需要一个资源可用这个函数创建,这个函数的实现超出了pool包的范围, 并且需要由包的使用者实现
	closed    bool                      // 表示Pool是否已经被关闭
}

// ErrPoolClosed 表示请求了一个已经关闭的池
var ErrPoolClosed = errors.New("Pool has been closed")

// New 创建一个用来管理资源的池
// 这个池需要一个可以分配新资源的函数, 并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire: ", "New Resource")
		return p.factory() // 创建一个资源
	}
}

// Release 将一个使用后的资源放在池里
func (p *Pool) Release(r io.Closer) {
	// 保证本操作和Close操作的安全
	p.m.Lock() // 和Close方法中的互斥量是同一个,这样可以阻止这两个方法在不同goroutine里同时运行
	defer p.m.Unlock()

	// 如果池已经被关闭, 销毁
	if p.closed {
		r.Close() // 池已关闭,无法将资源重新放回, 直接销毁它
		return
	}

	select {
	// 试图将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:", "Closing")
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close 会让资源池停止工作, 并关闭所有现有的资源
func (p *Pool) Close() {
	// 保证本操作与Release操作的安全
	p.m.Lock()
	defer p.m.Unlock() // 互斥量加锁,并在函数返回时解锁

	// 如果pool已被关闭, 什么都不做
	if p.closed {
		return
	}

	// 将池关闭
	p.closed = true

	// 在清空通道里的资源之前,将通道关闭
	// 如果不这样做, 会发生死锁
	close(p.resources)

	// 关闭资源
	for r := range p.resources {
		r.Close()
	}
}

const (
	maxGoroutines   = 25 // 要使用的goroutine的数量
	pooledResources = 2  // 池中的资源的数量
)

// dbConnection 模拟要共享的资源
type dbConnection struct {
	ID int32
}

// 实现io.Close接口, 以便dbConnection可以被池管理.
// Close可以用来完成任意资源的释放管理
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter用来给每个连接分配一个unique id
var idCounter int32

// createConnection是一个工厂,当需要一个新连接时,池会调用这个函数
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建用来管理连接的池
	p, err := New(createConnection, pooledResources)

	if err != nil {
		log.Println(err)
		return
	}

	// 使用池里的连接来完成查询
	for query := 0; query < maxGoroutines; query++ {
		// 每个goroutine需要自己复制一份要查询值的副本,不然所有的查询会共享同一个查询量
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("Shutdown Program.")
	p.Close()
}

func performQueries(query int, p *Pool) {
	// 从池里请求一个连接
	conn, err := p.Acquire()

	if err != nil {
		log.Println(err)
		return
	}

	// 将该连接释放回池里
	defer p.Release(conn)

	// 用等待模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
