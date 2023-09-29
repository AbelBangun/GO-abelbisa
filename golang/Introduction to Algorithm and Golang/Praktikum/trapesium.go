package main

import "fmt"

func main() {

	var s1 int
	var s2 int
	var t int

	fmt.Println("masukkan nilai sisi 1: ")
	fmt.Scan(&s1)

	fmt.Println("masukkan nilai sisi 2: ")
	fmt.Scan(&s2)

	fmt.Println("masukkan nilai tinggi: ")
	fmt.Scan(&t)

	L := ((s1 + s2) * t) / 2
	fmt.Println(L)
}
