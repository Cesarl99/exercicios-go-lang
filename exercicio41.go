package main

import (
	"fmt"
)

func main() {
	a := i()
	fmt.Println(a())

}
func i() func() int {
	x := 0
	return func() int {
		resultado := x + 10
		return resultado
	}
}

/*
- Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, onde esta outra função
faz uso de uma variável alem de seu scope.
*/
