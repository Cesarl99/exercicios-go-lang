package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	///fmt.Println(users)
	///utilização de encoder
	/*enco := json.NewEncoder(os.Stdout) ///1-criamos um encoder
		err := enco.Encode(users)          ///2-conseguimos acessar o encorder criado,para mostrar na tela
		if err != nil {
			fmt.Println("erro", err)

		}
	--------------------------------------------------------------------------------------------------
		PODE SER ESCRITO DA MANIERA ACIMA QUANDO NA MANEIRA ABAIXO*/
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("erro", err)
	}

}

/*
	 Partindo do código abaixo, utilize NewEncoder() e Encode() para enviar as informações como JSON para Stdout.
    - https://play.golang.org/p/BVRZTdlUZ_
*/
