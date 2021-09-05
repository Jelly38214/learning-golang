#### Typo

| chapter          | where | modify                      |
| ---------------- | ----- | --------------------------- |
| 02.6             | L52   | +fmt.Println()              |
| 03.3.4 copy 函数 | L57   | -哥 +个                     |
| 03.4.1           | L25   | -assiment +assignment       |
| 03.5.1           | L101  | ` 0`` => `0`/ p2_0 =>`p2_0` |
| 03.6             | L42   | -平凡 +平方                 |
| 03.6             | L57   | -memory +value              |
| 03.7             | L1    | -加息 +解析                 |
| 04.2             | L5    | -重+中                      |
| 04.2             | L6    | -sStructure +aStructure     |
| 04.2             | L7    | -astructure +aStructure     |
| 04.3             | L3    | -重 +中                     |
| 04.3             | L12   | -并没有                     |

Go 内置包

- time: 处理时间与日期;
  - 15 表示小时，04 表示分钟，05 表示秒，PM 转大写，pm 转小写
  - Jan 解析月份, 2006 解析年，02 解析天，Mon 解析周几
- strings: 处理字符串

Go 全局函数

- len
- cap
- append
- copy

特殊函数

- iota: 常量生成器

> 指针: `*`取值操作符, `&`取址操作符

Go 中的方法是一种特殊类型的函数,唯一的区别在于,必须在函数名前面加入一个额外的参数:`接收方`
场景:分组函数并将其绑定到自定义类型,通过方法将行为添加到结构(类似 JS 的对象的方法,该方法能使用这个对象内的属性), 这个方法能使用这个结构的其他字段
语法

```golang
  func (receiver type) MethodName(parameters...) {
    // Method functionality
  }
```

大写标识符 => 公开方法
非大写标识符 => 私有方法
Go 中的封装仅在程序包之间有效,换句话说,你只能隐藏来自其他程序包的`实现详细信息`, 而不能隐藏程序包本身

Go中的接口是一种抽象类型,只包括具体类型必须拥有或实现的方法,类似蓝图
Go没有实现接口的关键字.当Go中的接口具有接口所需的所有方法时,则满足`按类型的隐式实现`

Go有两种类型的包: 
 - Executable(main package is used to make an executable type package.): Generates a file that we can run
 - Reusable: code used as 'helpers'. Good place to put reusable logic


[Go的随机数](https://juejin.cn/post/6844903959065264141):
  在Go中获取的随机数,是基于某一个种子(seed)来的.也就是说,如果seed相同,那么获取的随机数就会相同. 解决这个现象的方法有两种:
    - 整个程序全局只初始化调用一次seed
    - 每次使用纳秒级别的种子

  其中第二种官方不推荐,因为在高并发下,及时是纳秒级 ,也可能获得相同的纳秒

类型转换:将某个类型转换成指定类型
语法:Type we want(Value we have)
例子: []byte("Hi there!") // 将字符串转成字节切片


new与make的区别
1. 两者都是用来做内存分配的(创建类型实例)
2. make只用于slice, map, channel的初始化, 返回类型本身
3. new用于类型的内存分配, 并且内存对应的值为类型对应的零值, 返回类型的指针

计算底层数组容量为k的切片slice[i:j]的长度和容量
- 长度: j - i
- 容量: k - i

方法: 方法能给用户定义的类型添加新的行为,实际上也是函数,只是在声明时,在关键字func和方法名之间增加一个接收者(receiver)

结构体和slice和数组的CRUD,掌握


go select有无default语句的区别, select/case语句