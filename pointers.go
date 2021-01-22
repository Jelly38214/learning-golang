package main

import "fmt"

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {
	i := -10
	j := 25

	pI := &i
	pJ := &j

	fmt.Println("pI memory: ", pI)
	fmt.Println("pJ memory: ", pJ)
	fmt.Println("pI value: ", *pJ)
	fmt.Println("pJ value: ", *pJ)

	fmt.Println()

	*pI = 123456
	*pI--
	fmt.Println("i: ", i)

	fmt.Println()

	getPointer(pJ)
	fmt.Println("j:", j) // => 625
	k := returnPointer(12)
	fmt.Println(*k) // => 144
	fmt.Println(k)
}
