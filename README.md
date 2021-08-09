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
