package main

import (
	"fmt"
)

var x [5]int

func main() {
	x[0] = 1
	x[1] = 5
	x[2] = 7
	x[3] = 7
	x[4] = 9
	for indice, valor := range x {
		fmt.Println(indice, valor)
		fmt.Printf(" ")
	}
	fmt.Printf("%T", x)
}

/*
- Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.

*/
