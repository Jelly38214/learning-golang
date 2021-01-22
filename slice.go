package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func printSlice(x []int) {
	for _, number := range x {
		fmt.Println(number, " ")
	}
}

func sortSlice(myslice []aStructure) {
	sort.Slice(myslice, func(i, j int) bool {
		return myslice[i].height < myslice[j].height
	})

	fmt.Println(":<", myslice)

	sort.Slice(myslice, func(i, j int) bool {
		return myslice[i].height > myslice[j].height
	})

	fmt.Println(">:", myslice)
}

func main() {
	aSlice := []int{-1, 0, 4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)

	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100) // double cap of slice
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Lenght: %d\n", cap(aSlice), len(aSlice))

	fmt.Println("\n=============\n")

	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}

	fmt.Print("a6:", a6)
	fmt.Print("a4:", a4)
	fmt.Println()

	copy(a6, a4) // copy a4 into a6 that is larger than a4
	fmt.Print("a6:", a6)
	fmt.Print("a4:", a4)

	fmt.Println("\n=============\n")

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}

	fmt.Print("b6: ", b6)
	fmt.Print("b4: ", b4)
	fmt.Println()

	copy(b4, b6) // copy b6 into b4 that is less than b4
	fmt.Print("b6: ", b6)
	fmt.Print("b4: ", b4)
	fmt.Println()

	fmt.Println("\n=============\n")

	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, -1, 1, -1, -5, 5}

	copy(s6, array4[0:])
	fmt.Print("array4:", array4[0:])
	fmt.Print("s6:", s6)
	fmt.Println()

	fmt.Println("\n=============\n")

	array5 := [5]int{5, -5, 5, -5, 5}
	s7 := []int{7, 7, -7, 7, -7, 7, 7}

	copy(array5[0:], s7)
	fmt.Print("array5:", array5)
	fmt.Print("s7: ", s7)
	fmt.Println()

	fmt.Println("\n=============\n")
	sortSlice([]aStructure{aStructure{person: "A", height: 10, weight: 20}, aStructure{person: "B", height: 30, weight: 40}})
}
