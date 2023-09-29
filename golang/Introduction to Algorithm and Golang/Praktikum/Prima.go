package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		b := 0
		for j := 1; j < 10; j++ {
			if i%j == 0 {
				b++
			}
		}
		if b == 2 && b != 1 {
			fmt.Println(i)
		}
	}
}
