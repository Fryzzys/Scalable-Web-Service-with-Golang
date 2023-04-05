package main

import "fmt"

func main() {
	perulangan() // memanggil function perulangan
}

func perulangan() {
	for i := 0; i < 5; { // melakukan perulangan variable i sebanyak 5 kali
		fmt.Printf("Nilai i : %d\n", i) // menampilkan string dan nilai variable i
		i++
		if i == 5 {
			for j := 0; j <= 10; j++ { // melakukan perulangan variable j sebanyak 10 kali
				if j == 5 {
					for i, char := range "САШАРВО" { // melakukan perulangan untuk menampilkan string
						fmt.Printf("character %#U starts at byte position %d\n", char, i) // menampilkan string berupa unicode format dan character nya
					}
					continue // memaksa untuk maju ke perulangan kedepanya
				} else {
					fmt.Printf("Nilai j : %d\n", j) // menampilkan string dan nilai variable j
				}
			}
		}
	}
}
