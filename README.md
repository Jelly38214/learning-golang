#### Typo

| chapter | where | modify         |
| ------- | ----- | -------------- |
| 02.6    | L52   | +fmt.Println() |
| 03.3.4 copy函数    | L57   | -哥 +个 |
| 03.4.1    | L25   | -assiment +assignment |
| 03.5.1   | L101   | `0`` => `0` / p2_0 => `p2_0`|
| 03.6    | L42   | -平凡 +平方 |
| 03.6    | L57   | -memory +value |
| 03.7    | L1   | -加息 +解析 |


Go内置包
  - time: 处理时间与日期; 
    - 15表示小时，04表示分钟，05表示秒，PM转大写，pm转小写
    - Jan解析月份, 2006解析年，02解析天，Mon解析周几
  - strings: 处理字符串

Go全局函数
  - len
  - cap
  - append
  - copy

特殊函数
  - iota: 常量生成器

> 指针: `*`取值操作符, `&`取址操作符
