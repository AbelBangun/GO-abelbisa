package main

import "fmt"

func main() {
	segitiga()
}

func segitiga() {

	for i := 1; i <= 5; i++ {
		for y := 5; y >= i; y-- {
			fmt.Print(" ")
		}
		for z := 1; z <= i; z++ {
			fmt.Print("* ")
		}
		fmt.Println(" ")
	}
}
