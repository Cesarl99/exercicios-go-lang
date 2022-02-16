package main

import (
	"fmt"
)

func main() {
	for x := 65; x < 91; x++ {
		println(x)
		for y := 0; y < 3; y++ {
			fmt.Printf("%c", x)
		}
		println(" ")
	}

}
