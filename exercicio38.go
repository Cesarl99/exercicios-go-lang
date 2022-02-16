package main

import (
	"fmt"
)

func main() {
	y := 89

	z := func(y int) {
		fmt.Println(y, "vezes 250 é:")
		fmt.Println(y * 250)
	}
	z(y)
}

/*
- Atribua uma função a uma variável.
- Chame essa função.

*/
