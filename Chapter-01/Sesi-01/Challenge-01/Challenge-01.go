package main

import "fmt"

func main() {
	var i int = 21                // deklarasi variable i bertipe intiger
	var j bool = true             // deklarasi variable j bertipe bolean
	var k int = 15                // deklarasi variable k bertipe intiger
	var l float64 = 123.456       // deklarasi variable l bertipe float
	var uniCode1 int32 = '\u042F' // deklarasi variable UniCode1 bertipe int32
	var uniCode2 int32 = 'Я'      // deklarasi variable UniCode2 bertipe int32

	fmt.Printf("%v\n", i)        // menampilkan nilai variable i = 21
	fmt.Printf("%T\n", i)        // menampilkan tipe data int
	fmt.Printf("%%\n")           // menampilkan tanda %
	fmt.Printf("%t\n", j)        // menampilkan nilai bolean true
	fmt.Printf("%c\n", uniCode1) // menampilkan unicode Rusia = Я (ya)
	fmt.Printf("%b\n", i)        // menampilkan nilai biner dari 21
	fmt.Printf("%d\n", i)        // menampilkan nilai base 10 = 21
	fmt.Printf("%o\n", i)        // menampilkan nilai base 8 = 25
	fmt.Printf("%x\n", k)        // menampilkan nilai base 16 = f
	fmt.Printf("%X\n", k)        // menampilkan nilai base 16 = f
	fmt.Printf("%U\n", uniCode2) // menampilkan unicode karakter Rusia Я
	fmt.Printf("%f\n", l)        // menampilkan nilai float 123.456000
	fmt.Printf("%E\n", l)        // menampilkan nilai float scientific 1.234560E+02
}
