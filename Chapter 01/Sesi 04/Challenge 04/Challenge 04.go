package main

import (
	"fmt"
	"os"
	"strconv"
)

type person struct { // membuat struct person yang di dalamnya terdapat beberapa data yang betipe string dan int
	nama      string
	umur      int
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	input, _ := strconv.Atoi(os.Args[1]) // deklarasi variable untuk menampung inputan dari CLI dengan maksimal argument yaitu 1, kemudian nilainya dirubah ke intiger
	var allStudentClass = []person{      // deklarasi variable dengan gabungan struct yang didalamnya terdapat data student
		{nama: "Agus Budiawan", umur: 30, alamat: "Jakarta", pekerjaan: "Buruh", alasan: "Ingin belajar pemrograman GO"},
		{nama: "Lili Astuti", umur: 28, alamat: "Magelang", pekerjaan: "Penjahit", alasan: "Ingin belajar pemrograman GO"},
		{nama: "Brahma", umur: 39, alamat: "Yogyakarta", pekerjaan: "Petani", alasan: "Ingin belajar pemrograman GO"},
		{nama: "Eliyah Firmansyah", umur: 29, alamat: "Sulawesi", pekerjaan: "Buruh", alasan: "Ingin belajar pemrograman GO"},
		{nama: "Pirman Bin", umur: 25, alamat: "Kalimantan", pekerjaan: "Fotografer", alasan: "Ingin belajar pemrograman GO"},
		{nama: "Iskandar", umur: 26, alamat: "Semarang", pekerjaan: "Wiraswasta", alasan: "Ingin belajar pemrograman GO"},
		{nama: "Agus Riyadi", umur: 36, alamat: "Bali", pekerjaan: "PNS", alasan: "Ingin belajar pemrograman GO"},
		{nama: "Mislan", umur: 33, alamat: "Solo", pekerjaan: "PNS", alasan: "Ingin belajar pemrograman GO"},
		{nama: "Umi Alifah", umur: 27, alamat: "Mojokerto", pekerjaan: "Seniman", alasan: "Ingin belajar pemrograman GO"},
		{nama: "Santi Sardi", umur: 24, alamat: "Gorontalo", pekerjaan: "Wiraswasta", alasan: "Ingin belajar pemrograman GO"},
	}
	for i, student := range allStudentClass { // melakukan perulangan dan data allStudentClass di tampung kedalam var student
		num := input - 1 // digunakan untuk mengurangi nilai dari input / input - 1 (untuk menyamakan index dari student)
		if num == i {    // melakukan pengecekan jika var num (input dari CLI) == i (index dari var student) akan menampilkan data ke-i
			fmt.Printf("Nama: %s\nUmur: %d\nAlamat: %s\nPekerjaan: %s\nAlasan: %s", student.nama, student.umur, student.alamat, student.pekerjaan, student.alasan) // menampilkan output berupa string dari data student
		}
	}
}
