package main

import "fmt"

func main() {

	// Variables

	var i int = 21
	var j bool = i > 10
	var k float64 = 123.456
	l := 15
	m := rune("Я") 

	fmt.Println("Halo, ini adalah Hasil Code Chalenge 1 ( Advin Kautsar ) : ")
	fmt.Println("")

	// menampilkan nilai i
	fmt.Println("nilai i => ", i)
	// menampilkan tipe data nilai i
	fmt.Printf("Tipe data dari nilai i yaitu => %T \n", i)
	//menampilkan tanda %
	fmt.Printf("menampilkan tanda %%\n")
	// menampilkan nilai j bernilai true
	fmt.Println("Nilai J bernilai =>", j)

	fmt.Println("")

	//menampilkan base 2 dari nilai i
	fmt.Printf("ini merupakan base2 dari nilai i => %b\n", i)
	//menampilkan hurur Я
	fmt.Println("ini akan menampilkan =>", string(m))
	// menampilkan base10 dari nilai i
	fmt.Printf("ini merupakan base10 dari nilai i => %d\n", i)
	// menampilkan base8 dari nilai i
	fmt.Printf("ini merupakan base8 dari nilai i => %o\n", i)
	// menampilkan base16
	fmt.Printf("ini merupakan base16  => %x\n", l)
	// menampilkan base16
	fmt.Printf("ini merupakan base16  => %X\n", l)
	// menampilkan unicode
	fmt.Printf("ini merupakan unicode karakter Я => %U\n", m)

	fmt.Println("")

	//menampilkan float dari k
	fmt.Printf("ini merupakan float dari nilai k  => %.6f\n", k)
	//menampilkan float dari k
	fmt.Printf("ini merupakan float scientific dari k  => %e\n", k)

}
