package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	sorvete   []string
}

func main() {
	pessoa1 := pessoa{
		nome:      "cesar",
		sobrenome: "luis",
		sorvete:   []string{"chocolate", "creme", "napolitano"}}
	pessoa2 := pessoa{
		nome:      "jose",
		sobrenome: "silva",
		sorvete:   []string{"creme", "limão", "morango"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sorvete {
		fmt.Println("\t", v)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sorvete {
		fmt.Println("\t", v)
	}
}

/*
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando
range na slice que contem os sabores de sorvete.


*/
