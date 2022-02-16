package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (estrut1 sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", estrut1.lat, estrut1.long, estrut1.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		erroNovo := fmt.Errorf("está errado:%v", f)
		return 0, sqrtError{"3236", "454512", erroNovo}
	}
	return 42, nil
}
