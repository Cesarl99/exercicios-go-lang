package main

import (
	"fmt"
)

func main() {
	a := (10 == 20)
	b := (3 != 6)
	c := (9 <= 2)
	d := (11 < 3)
	e := (65 >= 36)
	f := (236 > 521)

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)

}

/*escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
- Demonstre estes valores.*/
