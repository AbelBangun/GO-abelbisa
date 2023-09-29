package main

import "fmt"

func main() {

	var angka int

	fmt.Println("masukkan angka: ")
	fmt.Scan(&angka)

	if angka%2 == 0 {
		fmt.Println("bilangan genap")
	} else if angka%2 != 0 {
		fmt.Println("bilangan ganjil")
	}

}
