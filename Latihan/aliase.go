package main

import "fmt"

func main() {
	//deklarasi tipe data alias dengan second
	//type nama_alias  = nama_tipe_data

	type second = uint

	var hour second = 3600

	fmt.Printf("hour type %T\n", hour) // =>hour type : uint
}
