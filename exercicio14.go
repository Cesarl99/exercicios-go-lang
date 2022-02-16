package main

import (
	"fmt"
)

func main() {
	anonasc := 1999
	anoatual := 2022
	for {
		if anonasc > anoatual {
			break
		}
		fmt.Println(anonasc)
		anonasc++
	}

}

/*- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/
