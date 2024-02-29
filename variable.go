package main

import "fmt"

func main() {
	var name string = "faris"
	var umur int = 23

	name = "rasya"
	umur = 22

	//variabel dalam go juga bisa dituliskan tanpa menginisiasi apa tipe dari variabelnya
	var hobi = "nonton film"
	var nohp = 10000000000

	//penamaan variabel juga bisa disingkat dengan tidak perlu menggunakan keyword var
	pacar := "faris"
	nohppacar := 00000000

	//dalam go juga bisa membuat lebih dari satu variabel pada satu line yang sama
	var student1, student2, student3 string = "a", "b", "c"
	var first, second, third int

	first, second, third = 1, 2, 3

	//jika ingin menuliskam dengan tipe data yang berbeda beda bisa dituliskan seperti biasa atau menggunakan short declaration
	var nama, age, address = "fioni", 22, "jakarta"

	//dalam go, tidak boleh ada variabel yang menggangur, jika ada,maka akan terjadi error. untuk mengatasinya bisa menggunakan underscore

	var firstVariable, secondVariable, thirdVariable string
	_, _, _ = firstVariable, secondVariable, thirdVariable

	fmt.Println("halo namaku adalah ", name)
	fmt.Println("umurku ", umur)
	fmt.Println("hobik adalah ", hobi)
	fmt.Println("nomer hp ", nohp)
	fmt.Println("nama pacar aku", pacar)
	fmt.Println("nomer hpnya", nohppacar)
	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)
	fmt.Println(nama, age, address)
	fmt.Printf("Variable nama merupakan tipe variable %T \n", name)
	fmt.Printf("halo namaku adalah %s, umurku %d dan hobiku %s", name, age, hobi)

}
