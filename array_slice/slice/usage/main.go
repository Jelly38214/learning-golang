package main

import "fmt"

// 对切片的赋值操作和数组完全一致
func main() {
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)

	slice[1] = 25
	fmt.Println(slice)

	// 基于老的切片创建一个新的切片
	slice1 := slice[1:3]
	fmt.Println(slice1)

	// 现在两个切片共享同一个底层数组,如果一个切片修改了该底层数组的共享部分,另一个切片也能感知到
	slice1[1] = 35
	fmt.Println(slice, slice1)

	// 使用append函数增大切片长度,append会返回一个新的切片
	slice2 := append(slice1, 60) // 当前容量够用,append操作将可用的元素合并到切片的长度,并对其进行赋值
	fmt.Println(slice, slice1, slice2)

	// 当前容量不够用,append会创建一个新的底层数组,将被引用的先用的值复制到新的数组(针对这个切片引用的值才复制到新的数组),再追加新的值
	// 此时slice3跟其他的slice不共享共同一个数组,它有自己的新的底层数组
	slice3 := append(slice2, 70, 80)
	fmt.Println(slice, slice1, slice2, slice3, cap(slice3))

	// 使用第三个索引参数限制切片容量
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 长度为1, 容量为2
	sliceOfSource := source[2:3:4] // source[i:j:k], k为从索引位置i,加上希望容量中包含的元素的个数
	fmt.Println(sliceOfSource)

}
