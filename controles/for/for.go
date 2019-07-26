package main

import (
	"fmt"
)

func main() {

	i := 1
	for i <= 10 { //não tem while em GO
		fmt.Println(i)
		i++
	}

	for i2 := 0; i2 <= 20; i2 += 2 {
		fmt.Printf("%d", i2)
	}

	fmt.Println("\nMisturando...")
	for i3 := 1; i3 <= 10; i3++ {
		if i3%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Ímpar")
		}
	}
}
