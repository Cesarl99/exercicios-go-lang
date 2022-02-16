package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go loop(c)
	for v := range c {
		fmt.Println(v)
	}
}
func loop(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

/*
- Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.
*/
