package main

import (
	"fmt"
)

func main() {
	for x := 1; x < 10001; x++ {
		fmt.Println(x)
		fmt.Println("-/")
	}
	println("FIM")
}
