package main

import (
	"fmt"
)

func normalLoop() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}
		if i == 95 {
			break
		}

		fmt.Print(i, " ")
	}
}

func whileLoop() {
	i := 10
	for {
		if i < 0 {
			break
		}

		fmt.Print(i, " ")
		i--
	}

	fmt.Print()
}

func doWhileLoop() {
	i := 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}
		fmt.Print(i, " ")
		i++
	}
}

func rangeLoop() {
	anArray := [5]int{4, 3, 2, 1, 0}
	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value)
	}
}

func main() {
	normalLoop()
	fmt.Println("\n========")
	whileLoop()
	fmt.Println("\n========")
	doWhileLoop()
	fmt.Println("\n========")
	rangeLoop()
}
