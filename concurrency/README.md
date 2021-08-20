传统安全同步访问共享资源的两种锁: 原子锁(atomic)/互斥锁(mutual exclusion)

原子函数和互斥锁都能工作(传统方式),但依靠它们都不会让编写并发程序变得更简单, 更不容易出错,或者有趣

> 在Go语言中,主要使用通道来发送和接受需要共享的资源

当一个资源需要在goroutine之间共享时,通道在goroutine之间架起了一个通道,并提供了确保同步交换数据的机制.

声明通道时,需要指定将要被共享的数据的`类型`
可以通过通道共享内置类型,命名类型,结构类型和引用类型的值或指针

> 创建通道需要使用内置函数make(syntax: make(chan type))
```go
  // 无缓冲的字符串通道
  unBuffered := make(chan int)

  // 有缓冲的字符串通道
  buffered = make(chan string, 10)
```

向通道读与写符号都是`<-`, 区别在于chan在该符号的位置
chan <- 数据// 向其写入数据
<- chan // 读其数据

有无缓冲区的通道的区别:
  - 
atomic 的 add/store/load 区别