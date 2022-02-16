package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type ordeneage []user

func (x ordeneage) Len() int           { return len(x) }
func (x ordeneage) Less(i, j int) bool { return x[i].Age < x[j].Age }
func (x ordeneage) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ordenalast []user

func (y ordenalast) Len() int           { return len(y) }
func (y ordenalast) Less(i, j int) bool { return y[i].Last < y[j].Last }
func (y ordenalast) Swap(i, j int)      { y[i], y[j] = y[j], y[i] }

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

	sort.Sort(ordeneage(users))
	println("ORDENADO POR IDADE:")
	for i, v := range users {
		fmt.Println(i, "IDADE:", v.Age)
		fmt.Println("NOME:", v.First)
		fmt.Println("SOBRENOME:", v.Last)
		fmt.Println("FALAS:", v.Sayings)
		fmt.Println(" ")

	}

	///fmt.Println("ordenado por idade: ", users)
	println("ORDENADO PELO ULTIMO NOME:")
	sort.Sort(ordenalast(users))
	for i, v := range users {
		fmt.Println(i, "NOME:", v.First)
		fmt.Println("SOBRENOME:", v.Last)
		fmt.Println("IDADE:", v.Age)
		fmt.Println("FALAS:", v.Sayings)
		fmt.Println(" ")

	}

	///fmt.Println("ordenado pelo ultimo nome", users)
	/*
		for _, v := range users {
			sort.Strings((v.Sayings))
		}
		fmt.Println("as falas do u1 em ordem alfabetica é :\n", u1.Sayings)
		fmt.Println("as falas do u2 em ordem alfabetica é :\n", u2.Sayings)
		fmt.Println("as falas do u3 em ordem alfabetica é :\n", u3.Sayings)*/
}

///- Partindo do código abaixo, ordene os []user por idade e sobrenome.
