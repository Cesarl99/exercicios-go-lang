package main

import (
	"fmt"
)

type veiculo struct {
	portas int
	cor    string
}
type caminhonete struct {
	veiculo
	quatrorodas bool
}
type sedan struct {
	veiculo
	modelodeluxo bool
}

func main() {
	carro1 := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "preto",
		},
		quatrorodas: true,
	}
	carro2 := sedan{
		veiculo: veiculo{
			portas: 2,
			cor:    "vermelho",
		},
		modelodeluxo: true,
	}
	fmt.Println(carro1)
	fmt.Println(carro2)
	fmt.Println(carro1.cor)
	fmt.Println(carro2.portas)
}
