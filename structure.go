package main

import "fmt"

type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

func createStructure(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}

	return &myStructure{n, s, h}
}

func retStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}

	return myStructure{n, s, h}
}

func main() {
	type XYZ struct {
		X int
		Y int
		Z int
	}

	var s1 XYZ
	fmt.Println(s1.Y, s1.Z)

	p1 := XYZ{23, 12, -2}
	p2 := XYZ{Z: 12, Y: 13}

	fmt.Println(p1)
	fmt.Println(p2)

	pSlice := [4]XYZ{}
	pSlice[2] = p1
	pSlice[0] = p2

	// 结构体分配给数组以深拷贝的方式， 因此p1, p2的变动不会影响到在数组中的结构体
	p1.X = 100
	p2 = XYZ{1, 2, 3}

	fmt.Println(pSlice)

	/**
	* 结构体指针
	 */

	sp1 := createStructure("Mihalis", "Tsoukkalos", 123)
	sp2 := retStructure("Mihalis", "Tsoukkalos", 123)

	fmt.Println((*sp1).Name)
	fmt.Println(sp2.Name)
	fmt.Println(sp1)
	fmt.Println(sp2)
}
