package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// como não vou usar a variável i, coloco o _
	for _, num := range numeros {
		fmt.Println(num)
	}

	// acesando apenas o índice e não o número
	for num := range numeros {
		fmt.Println(num)
	}
}
