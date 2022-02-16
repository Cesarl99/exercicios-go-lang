package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	x = 42
	y = "James Bond"
	z = false
	s := fmt.Sprintf("%v,\n%v,\n%v", x, y, z)
	fmt.Print(s)

}

/* Utilizando a solução do exercício anterior:
   1. Em package-level scope, atribua os seguintes valores às variáveis:
       1. para "x" atribua 42
       2. para "y" atribua "James Bond"
       3. para "z" atribua false
   2. Na função main:
       1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
       2. Demonstre a variável "s".*/
