package main

import "fmt"

func main() {
	var buah = []string{"apel", "pisang", "duren"}

	for i := 0; i < len(buah); i++ {
		fmt.Printf("indeks ke %d, buahnya %s \n", i, buah[i])
	}

	//cara kedua membuat slice
	var hewan = make([]string, 4)

	_ = hewan

	hewan[0] = "hewan"
	hewan[1] = "anjing"
	hewan[2] = "babi"
	hewan[3] = "kucing"

	fmt.Println(hewan)

	//append func with ellipsis

	var hewan1 = []string{"anjing", "kucing", "burung"}
	var hewan2 = []string{"babi", "lalat", "jangkrik"}

	hewan1 = append(hewan1, hewan2...)

	fmt.Println(hewan1)

	fmt.Println("=======================================")
	var fruit1 = []string{"nanas", "semangka", "melon"}
	var fruit2 = []string{"apel", "manggis", "sirkaya"}

	nn := copy(fruit1, fruit2)

	fmt.Println(fruit2)
	fmt.Println(fruit1)
	fmt.Println(nn)

	var hewan3 = hewan[1:4]
	hewan4 := append(hewan[:2], "serigala")
	fmt.Println(hewan3)
	fmt.Println(hewan4)

}
