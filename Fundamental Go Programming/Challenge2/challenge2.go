package main

import "fmt"

func main() {

	// Variables
	char := []rune{'С', 'А', 'Ш', 'А', 'Р', 'В', 'О'}
	byte := 0

	fmt.Println("Halo, ini adalah Hasil Code Chalenge 2 ( Advin Kautsar ) : ")
	fmt.Println("")

	for i := 0; i <= 4; i++ {
		fmt.Println("Nilai i = ", i)
	}

	for j := 0; j <= 4; j++ {
		fmt.Println("Nilai j = ", j)
		if j == 4 {
			for _, huruf := range char {
				fmt.Printf("character : %U '%c' starts at by position %d \n", huruf, huruf, byte)
				byte += 2
			}
		}
	}

	for j := 6; j <= 10; j++ {
		fmt.Println("Nilai j =", j)
	}

}
