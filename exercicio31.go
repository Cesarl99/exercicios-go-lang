package main

import (
	"fmt"
)

func main() {

	x := struct {
		nome   string
		idade  int
		altura float64
		hobby  []string          ///slice
		comida map[string]string ///map
	}{
		nome:   "luis",
		idade:  25,
		altura: 1.70,
		hobby:  []string{"futebol", "video game", "luta"}, ///slice
		comida: map[string]string{ ///map
			"café da manhã": "pão",
			"almoço":        "arroz",
			"janta":         "bife"},
	}
	fmt.Println(x)
}
