package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//tipos inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)
	var b byte = 255
	fmt.Println("o byte é", reflect.TypeOf(b))

	//com sinal
	i1 := math.MaxFloat64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 = 'a' // representa um mapeamento da tabela Unicode
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//números reais
	var x float32 = 49.99
	fmt.Println("o tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo) //negação

	//string
	s1 := "Olá meu nome é Helio"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//string com múltiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Hélio`
	fmt.Println("O tamanho da string é", len(s2))

	//caracter único (equivale ao char em outras linguagens)
	c1 := 'a'
	fmt.Println("O tipo de caracter único é", reflect.TypeOf(c1))

}
