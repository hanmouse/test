package main

import "fmt"

type Suite int

const (
	Spades Suite = iota
	Hearts
	Diamonds
	Clubs
)

func main() {
	var s Suite = Hearts
	fmt.Print(s)
	switch s {
	case Spades:
		fmt.Println(" are best.")
	case Hearts:
		fmt.Println(" are second best.")
	default:
		fmt.Println(" aren't very good.")
	}
}

func (s Suite) String() string {
	return [...]string{"Spades", "Hearts", "Diamonds", "Clubs"}[s]
}
