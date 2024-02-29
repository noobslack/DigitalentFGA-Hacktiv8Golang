package main

import "fmt"

func main() {
	//const merupakan sebuah deklarasi variabel yang nilainya tidak dapat diubah
	const full_name string = "faris"
	fmt.Println("namaku adalah", full_name)

	//operator
	var value = (2 + 3) * 3
	fmt.Println(value)

	//relational operator
	var firstCondition bool = 1 <= 3
	var secondCondition bool = 1 > 9
	var thirdCondition bool = 10 != 3

	fmt.Println(firstCondition)
	fmt.Println(secondCondition)
	fmt.Println(thirdCondition)	

	var wrong = false
	var right = true

	var wrongAndright = wrong && right
	fmt.Printf("\t(%t)", wrongAndright)

}
