package main

import "fmt"

// pada go hanya terdapat satu perulangan, yaitu hanya dengan menggunakan for
func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("angka ke ", i)
	// }

	// //ada cara lain, mirip dengan penggunaan while
	// j := 0
	// for j < 3 {
	// 	fmt.Println(j)
	// 	j++
	// }

	// //cara ketiga, dengan tidak memberikan pengkondisian apapun

	// var k = 0

	// for {
	// 	fmt.Println(k)

	// 	k++
	// 	if k == 5 {
	// 		break
	// 	}
	// }

	// implementasi untuk mencari bilangan ganjil dan genap
	for angka := 1; angka <= 10; angka++ {
		if angka%2 == 1 {
			continue
		}
		if angka >= 10 {
			break
		}
		fmt.Println(angka)
	}

	//nested loop
outerloop:
	for i := 0; i <= 5; i++ {
		fmt.Println("perulangan ke ", i)
		for j := i; j < 5; j++ {
			if i == 2 {
				break outerloop
			}
			fmt.Print(j, "")
		}

		fmt.Println()
	}
	fmt.Println("========================")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

}
