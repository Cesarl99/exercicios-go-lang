package main

import (
	"fmt"
)

func main() {
	fmt.Println(inteiro())
	fmt.Println(stint())
}
func inteiro() int {
	return 10
}
func stint() (string, int) {
	return "vinte", 20
}

/*
- Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados

*/
