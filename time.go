package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("Epoch Time:", time.Now().Unix())
	t := time.Now()

	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time different:", t1.Sub(t))

	fmt.Println()

	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)
	loc, _ := time.LoadLocation("Europe/Paris")
	LondonTime := t.In(loc)
	fmt.Println("Pairs:", LondonTime)

	fmt.Println()

	var myTime string

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s string\n", filepath.Base(os.Args[0]))
		os.Exit(0)
	}

	myTime = os.Args[1]

	d, err := time.Parse("15:04", myTime)
	if err == nil {
		fmt.Println("Full", d)
		fmt.Println("Time", d.Hour(), d.Minute())
	} else {
		fmt.Println(err)
	}

}
