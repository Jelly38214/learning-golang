package main

import "fmt"

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			// This anonymous has not passed parameter i, So it'll get i of d2 scope which is zero.
			fmt.Print(i, " ")
		}()
	}
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	d1()
	fmt.Println()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
