package main

import "fmt"

func main() {

	// var score = 8

	// switch score {
	// case 8:
	// 	fmt.Println("perfect")
	// case 7:
	// 	fmt.Println("awesome")
	// }

	//bisa juga dilakukan perbandingan dengan menggunakan operator

	var nilai = 2

	if nilai > 6 {
		switch nilai {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")

		}

	} else {
		if nilai == 5 {
			fmt.Println("jelek")
		} else if nilai == 3 {
			fmt.Println("jelek bgt")
		} else {
			fmt.Println("ulangi lagi laa")
		}
	}

}
