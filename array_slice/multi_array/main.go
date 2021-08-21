package main

import "fmt"

/* 多维数组很容易管理具有父子关系的数据或者与坐标相关联的数据 */

func main() {

	// 使用数组字面量来声明并初始化一个二维整型数组
	array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array)

	array[1][1] = 22
	fmt.Println(array)

	/* 多维数组之间赋值 */
	var (
		array1, array2 [2][2]int
	)

	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40

	// 将array2的值复制给array1
	array1 = array2
	fmt.Println("array1: ", array1)

	// 每个数组都是一个值,可以仅复制其某一部分
	var array3 [2]int = array1[1]
	fmt.Println("array3: ", array3)

	var value30 int = array1[1][0]
	fmt.Println("value30: ", value30)

}
