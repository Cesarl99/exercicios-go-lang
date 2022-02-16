package main

import (
	"fmt"
)

func main() {
	x := 16
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}

/*- Crie um programa que:
  - Atribua um valor int a uma variável
  - Demonstre este valor em decimal, binário e hexadecimal*/
