package main

import (
	"fmt"
)

func main() {
	função2(função1)
}
func função1() {
	fmt.Println("inicio.....")
}
func função2(x func()) {
	fmt.Println("......final")
	x()
}

/*
- Callback: passe uma função como argumento a outra função.
*/
