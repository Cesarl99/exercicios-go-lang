package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Seu sistema Operacional é :", runtime.GOOS)
	fmt.Println("Seu ARCH é :", runtime.GOARCH)
}

/*
- Crie um programa que demonstra seu OS e ARCH.
- Rode-o com os seguintes comandos:
    - go run
    - go build
    - go install
*/
