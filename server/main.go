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

	players, _ := runner.ReadConfig()

	// TODO: load or generate map
	game.Map = &Map{
		Radius:  WorldRadius,
		Walls:   []*Wall{},
		Items:   []*Item{},
		Players: []*Player{},
	}
	game.SpawnPlayers(players)
	game.SendMapToObserver()
	game.GreetPlayers()

	for game.ShouldContinue() {
		for _, player := range game.Map.Players {
			game.DoTurn(player)
		}
		game.Tick()
		game.SendStateToObserver()
		game.Turn++
	}

	game.Runner.Log("that's all folks!")
	// TODO: scores
	game.Runner.End()
}
