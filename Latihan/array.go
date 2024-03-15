package main

import "fmt"

func main() {
	array := [4]int{1, 2, 3, 4}
	array[0] = 100

	for i, v := range array {
		fmt.Printf("indeks ke %d, value: %d\n", i, v)
	}

	//cara keuda
	for i := 0; i < len(array); i++ {
		fmt.Printf("indeks ke %d, value: %d\n", i, array[i])

	}

	//array multidimensional
	matrix := [2][3]int{{2, 3, 4}, {4, 5, 6}}
	for _, arr := range matrix {
		for _, value := range arr {
			fmt.Printf("%d", value)
		}
		fmt.Println()
	}

}
