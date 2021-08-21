package main

import "fmt"

// 迭代切片
func main() {
	slice := []int{10, 20, 30, 40}

	// 第一个值是当前迭代到的索引位置,第二个值是该位置对应元素的副本
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	// 注意返回的是副本,而不是元素的引用
	// 可以发现Value-Addr不等于ElementAddr
	// &value1总是相同的,因为它是本次迭代过程中返回的新的变量,在迭代过程中根据切片依次被赋值
	for index1, value1 := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElementAddr: %X\n", value1, &value1, &slice[index1])
	}
}
