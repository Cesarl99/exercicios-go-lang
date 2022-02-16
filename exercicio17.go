package main

import (
	"fmt"
)

func main() {
	esportefavorito := "basquete"
	switch esportefavorito {
	case "futebol":
		fmt.Println("seu esporte favorito é futebol")
	case "volei":
		fmt.Println("seu esporte favorito é volei")
	case "basquete":
		fmt.Println("seu esporte favorito é basquete")

	}
}

/*
- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do
 tipo string com identificador "esporteFavorito".
- Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/
