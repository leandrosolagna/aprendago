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

type ordenaIdade []user

func (x ordenaIdade) Len() int           { return len(x) }
func (x ordenaIdade) Less(i, j int) bool { return x[i].Age < x[j].Age }
func (a ordenaIdade) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ordenaSobrenome []user

func (x ordenaSobrenome) Len() int           { return len(x) }
func (x ordenaSobrenome) Less(i, j int) bool { return x[i].Last < x[j].Last }
func (a ordenaSobrenome) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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

	fmt.Println("Desordenado:", users)
	fmt.Println()

	sort.Sort(ordenaIdade(users))
	fmt.Println("Idade:\n", users)

	sort.Sort(ordenaSobrenome(users))
	fmt.Println("Sobrenome:\n", users)

	fmt.Println()

	fmt.Println("Por dizeres:")
	for i, v := range users {
		fmt.Println(i, "\tIdade:", v.Age, "\tNome completo:", v.First, v.Last)
		for _, c := range v.Sayings {
			fmt.Println("\t", c)
		}
		fmt.Println()
	}
}
