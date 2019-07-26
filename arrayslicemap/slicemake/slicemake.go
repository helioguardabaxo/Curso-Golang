package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // 10 elementos e 20 elementos totais
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// adicionado elementos a mais expande a capacidade do slice
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
