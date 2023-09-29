package main

import "fmt"

func main() {

	var angka int
	fmt.Println("masukkan angka: ")
	fmt.Scan(&angka)

	if angka >= 80 && angka <= 100 {
		fmt.Println("A")
	}

	if angka >= 65 && angka <= 79 {
		fmt.Println("B")
	}
	if angka >= 50 && angka <= 64 {
		fmt.Println("C")
	}
	if angka >= 35 && angka <= 49 {
		fmt.Println("D")
	}
	if angka > 0 && angka <= 34 {
		fmt.Println("E")
	}
	if angka <= 0 {
		fmt.Println("Nilai Invalid")
	}
}
