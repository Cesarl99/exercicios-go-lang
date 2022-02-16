package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "mark",
		Age:   25,
	}

	u2 := user{
		First: "jonh",
		Age:   36,
	}

	u3 := user{
		First: "carl",
		Age:   74,
	}
	users := []user{u1, u2, u3}
	fmt.Println(users)

	x, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(x))
}

/*
- Partindo do código abaixo, utilize marshal para transformar []user em JSON.
    - https://play.golang.org/p/U0jea43X55

*/
