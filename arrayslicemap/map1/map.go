package main

import (
	"fmt"
)

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[90923233] = "Pedro"
	aprovados[12456312] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12456312])
	delete(aprovados, 12456312)
	fmt.Printf(aprovados[12456312])

}
