package main

import (
	"fmt"
)

func main() {
	r := 25.025
	circulo(r)
	l := 3
	quadrado(l)
}
func circulo(r float64) {
	fmt.Println("o valor da area de circulo é", 2*3.14*r)
	///fmt.Println(2 * 3.14 * r)
}
func quadrado(l int) {
	fmt.Println("o valor da area do quadrado é ", l*l)
	///fmt.Println(l * l)
}

/