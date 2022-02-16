package main

import (
	"fmt"
)

func main() {
	slice := []int{36, 55, 63, 52, 100, 14, 1, 2, 5, 9}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", slice)
}

/*
 Usando uma literal composta:
    - Crie uma slice de tipo int
    - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.
*/
