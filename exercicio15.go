package main

import (
	"fmt"
)

func main() {
	for x := 10; x < 100; x++ {
		r := x % 4
		fmt.Printf("o resto de %v dividido por 4 é %v", x, r)
		println(" ")
	}

}

/*- Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
 */
