package main

import "fmt"

func main() {
	var currentYear = 2017

	if age := currentYear - 1998; age < 17 {
		fmt.Println("Kamu masih underage")
	} else {
		fmt.Println("kamu cukup umur")
	}
}
