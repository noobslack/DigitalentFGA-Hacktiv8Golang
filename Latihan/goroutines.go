package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("main execution started")

	go firstProcess(8)

	secondProcess(8)

	fmt.Println("No. of Goroutines", runtime.NumGoroutine())

	fmt.Println("main excecution is needed")
}

func firstProcess(index int) {
	fmt.Println("first process started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("first process func ended")
}

func secondProcess(index int) {
	fmt.Println("second process func ended")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("second process ended")
}
