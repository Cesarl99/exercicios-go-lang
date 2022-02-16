package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for x := 0; x < 10; x++ {
				canal <- x
			}
		}()
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-canal)
	}
}

/*
- Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
- Tire estes números do canal e demonstre-os.
*/
