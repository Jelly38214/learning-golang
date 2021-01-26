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
| 04.2             | L6    | -sStructure +aStructure      |
| 04.2             | L7    | -astructure +aStructure      |
| 04.3             | L3    | -重 +中      |
| 04.3             | L12    | -并没有     |

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
