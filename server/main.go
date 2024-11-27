package main

import (
	"fmt"
	"github.com/trojsten/ksp-proboj/client"
	"time"
)

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
		game.Runner.Log(fmt.Sprintf("TURN %d", game.Turn))
		for _, player := range game.Map.Players {
			start := time.Now()
			game.DoTurn(player)
			end := time.Now()
			if end.Sub(start) > time.Second {
				runner.KillPlayer(player.Name)
			}
		}
		game.Tick()
		game.SendStateToObserver()
		game.Turn++
	}

	game.Runner.Log("that's all folks!")
	game.SendScores()
	game.Runner.End()
}
