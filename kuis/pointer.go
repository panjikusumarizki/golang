package kuis

import (
	"fmt"
)

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "Mario"}

	gamer.AddGame("Borderlands")
	gamer.AddGame("FIFA22")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
