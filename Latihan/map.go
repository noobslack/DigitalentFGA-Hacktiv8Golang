package main

import "fmt"

func main() {
	var person map[string]string //deklarasi

	person = map[string]string{} //inisialisasi

	person["name"] = "faris"
	person["age"] = "17"
	person["address"] = "Pemalang"

	fmt.Println("name: ", person["name"])
	fmt.Println("age: ", person["age"])
	fmt.Println("address: ", person["address"])

	//cara kedua

	var mahasiswa = map[string]string{
		"name":    "rasya",
		"age":     "22",
		"jurusan": "sastra indonesia",
	}
	for key, value := range mahasiswa {
		fmt.Println(key, ":", value)
	}

	delete(mahasiswa, "jurusan")

	for key, value := range mahasiswa {
		fmt.Println(key, ":", value)
	}

	value, exist := person["age"]

	if exist {
		fmt.Println(value)

	} else {
		fmt.Println("value doesnt exist")

	}

	delete(person, "age")

	value, exist = person["age"]

	if exist {
		fmt.Println(value)

	} else {
		fmt.Println("value has been deleted ")
	}

	fmt.Println("##########################")
	sliceMap()

}

func sliceMap() {
	var mahasiswa = []map[string]string{
		{"name": "faris", "age": "17"},
		{"name": "rasya", "age": "18"},
		{"name": "julian", "age": "19"},
	}

	for i, student := range mahasiswa {
		fmt.Printf("index : %d, name : %s, age : %s\n", i, student["name"], student["age"])
	}

}
