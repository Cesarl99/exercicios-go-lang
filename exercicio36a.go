package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}
type circulo struct {
	raio float64
}

func (q quadrado) areaq() {
	resultado := q.lado * q.lado
	fmt.Println("area do quadrado é", resultado)
}
func (c circulo) areac() {
	resultadocic := math.Pi() * 2 * c.raio
	fmt.Println(resultadocic)
}
func main() {
	q := quadrado{
		lado: 20.3,
	}
	c := circulo{
		raio: 36.6,
	}
	c.areac
	q.areaq

}
