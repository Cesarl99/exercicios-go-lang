package main

import (
	"fmt"
)

const (
	_ = 2022 + iota
	ano2
	ano3
	ano4
	ano5
)

func main() {

	fmt.Println(ano2, ano3, ano4, ano5)
}

/*- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.*/
