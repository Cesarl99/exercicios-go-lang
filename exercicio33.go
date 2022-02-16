package main

import (
	"fmt"
)

func main() {
	sliceint := []int{1, 6, 9, 2, 4, 77}
	sliceint2 := []int{1, 6, 9, 2, 4, 777}
	fmt.Println(funcao1(sliceint...))
	fmt.Println(funcao2(sliceint2))
}
func funcao1(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func funcao2(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

/*
- Crie uma função que receba um parâmetro variádico do tipo int e
retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e
retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
*/
