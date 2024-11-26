package main

import "github.com/trojsten/ksp-proboj/client"

type Shooting struct {
	Attacker *Player
	Target   *Player
}

type Game struct {
	Runner        *client.Runner
	Map           *Map
	Turn          int
	TurnShootings []Shooting
	TurnYaps      []string
}

func main() {
	runner := client.NewRunner()
	game := Game{
		Runner: &runner,
		Turn:   0,
	}

	players, mapName := runner.ReadConfig()

	game.Map = &Map{
		Radius:  WorldRadius,
		Walls:   []*Wall{},
		Items:   []*Item{},
		Players: []*Player{},
	}

	if mapName != "" {
		game.LoadMap(mapName)
	}

	game.SpawnItems()
	game.SpawnPlayers(players)
	game.SendMapToObserver()
	game.GreetPlayers()

	for game.ShouldContinue() {
		game.TurnShootings = []Shooting{}
		game.TurnYaps = []string{}
		for _, player := range game.Map.Players {
			game.DoTurn(player)
		}
		game.Tick()
		game.SendStateToObserver()
		game.Turn++
	}

	game.Runner.Log("that's all folks!")
	game.SendScores()
	game.Runner.End()
}
