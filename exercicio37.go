package main

import (
	"fmt"
)

func main() {
	x := 12
	func(int) {
		fmt.Println(x, "dividido por 2 é:")
		fmt.Println(x / 2)
	}(x)
}

/*
- Crie e utilize uma função anônima.
*/
