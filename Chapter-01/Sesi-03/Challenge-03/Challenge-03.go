package main

import (
	"fmt"
)

func main() {
	perulangan() // memanggil function perulangan
}

func perulangan() {
	input := "selamat malam"        // input awal
	tampung := make(map[string]int) // menggunakan map untuk menampung tiap huruf dan jumlahnya

	for i := 0; i < len(input); i++ { // melakukan perulangan sebanyak panjang variable input
		fmt.Println(string(input[i])) // menampilkan output dari variable input berupa string secara vertikal
	}

	for _, count := range input { // melakukan perlulangan elemen saja
		tampung[string(count)] = tampung[string(count)] + 1 // menambahkan data elemen dari variable input ke map tampung dan menambah index +1 setiap huruf yang terduplikasi
	}

	fmt.Println(tampung) // menampilkan output berupa string dan jumlah huruf yang terduplikasi
}
