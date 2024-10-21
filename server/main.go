package main

import "github.com/trojsten/ksp-proboj/client"

type Game struct {
	Runner *client.Runner
	Map    *Map
	Turn   int
}

func main() {
	runner := client.NewRunner()
	game := Game{
		Runner: &runner,
		Turn:   0,
	}

	// TODO: load or generate map
	game.GreetPlayers()

	for game.ShouldContinue() {
		for _, player := range game.Map.Players {
			game.DoTurn(player)
		}
	}
}
