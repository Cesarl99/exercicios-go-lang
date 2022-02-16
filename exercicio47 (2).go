package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}
	fmt.Println("a lista de string sem ordenar e a seguinte \n", xs)
	sort.Strings(xs)
	fmt.Println("a lista de string ordenada é :\n", xs)
	fmt.Println("a lista de numeros int sem ordenar e a seguinte \n", xi)
	sort.Ints(xi)
	fmt.Println("a lista de nuemros int ordenada é :\n", xi)
}

/*
- Partindo do código abaixo, ordene a []int e a []string.
*/
