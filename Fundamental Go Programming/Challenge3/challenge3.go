package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("Halo, ini adalah Hasil Code Chalenge 3 ( Advin Kautsar ) : ")
	fmt.Println("")

	// variable
	inputan := "selamat malam"
	pecah := strings.Split(inputan, " ")

	for _, kata1 := range pecah {
		for _, kata2 := range kata1 {
			fmt.Println(string(kata2))
		} 
		fmt.Println()
	}

	output := make(map[string]int)
	for _, kata1 := range pecah {
		for _, kata2 := range kata1 {
			output[string(kata2)]++
		} 
	}

	fmt.Print("map[ ")
		for k,v := range output {
			fmt.Printf("%s:%d ", k, v)
		}
		fmt.Println("]")
	}