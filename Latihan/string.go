package main

import "fmt"

func main() {
	city := "faris"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index : %d, byte %d\n", i, city[i])
	}

	bytetoString()
}

func bytetoString() {

	name := []byte{102, 97, 114, 105, 115}

	fmt.Println(string(name))
}


