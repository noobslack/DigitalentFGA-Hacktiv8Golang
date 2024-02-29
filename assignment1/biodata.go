package main

import (
	"fmt"
	"os"
	"strconv"
)

type fgaStudent struct {
	no        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	allStudent := []fgaStudent{
		{no: 1, nama: "Faris", alamat: "Pemalang", pekerjaan: "Mahasiswa", alasan: "alasan 1"},
		{no: 2, nama: "Rasya", alamat: "Jatinangor", pekerjaan: "Frontend Engineer", alasan: "alasan 2"},
		{no: 3, nama: "Putri", alamat: "Bandung", pekerjaan: "UI/UX Designer", alasan: "alasan 3"},
		{no: 4, nama: "Zakaria", alamat: "Cianjur", pekerjaan: "QA Engineer", alasan: "alasan 4"},
		{no: 5, nama: "Tono", alamat: "Kuningan", pekerjaan: "Backend Engineer", alasan: "alasan 5"},
	}

	if len(os.Args) == 1 {
		fmt.Println("Tolong masukan nomer peserta")
	} else if len(os.Args) == 2 {
		search, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("input harus berupa angka", err)
			return
		}
		printStudent(allStudent, search)

	}

}

func printStudent(student []fgaStudent, search int) {
	foundValue := false

	for _, value := range student {
		if search == value.no {
			fmt.Println("Nama:", value.nama)
			fmt.Println("alamat:", value.alamat)
			fmt.Println("pekerjaan:", value.pekerjaan)
			fmt.Println("alasan:", value.alasan)
			foundValue = true
			break
		}

	}

	if !foundValue {
		fmt.Println("tidak ada peserta dengan nomor tersebut")
	}

}
