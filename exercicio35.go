package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (qualquerpessoa pessoa) mostrar() {
	fmt.Println(qualquerpessoa.nome, qualquerpessoa.sobrenome, qualquerpessoa.idade)
}
func main() {
	outrapessoa := pessoa{
		nome:      "cesar",
		sobrenome: "luis",
		idade:     22,
	}
	outrapessoa.mostrar()
}

/*
Crie um tipo struct "pessoa" que contenha os campos:
- nome
- sobrenome
- idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
*/
