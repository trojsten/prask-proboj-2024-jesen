package main

import (
	"fmt"
	"time"

	"github.com/trojsten/ksp-proboj/client"
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
			// Players are not paused at start, let's avoid unnecessary error messages in logs.
			if game.Turn != 0 {
				runner.ResumePlayer(player.Name)
			}
			start := time.Now()
			game.DoTurn(player)
			end := time.Now()
			//if end.Sub(start) > time.Second*10 {
			//	runner.KillPlayer(player.Name)
			//	game.Runner.Log(fmt.Sprintf("Killing player %s, turn took too long (%v)", player.Name, end.Sub(start).Seconds()))
			//}
			game.Runner.Log(fmt.Sprintf("PLAYER %s turn took %v ms", player.Name, end.Sub(start).Milliseconds()))

			runner.PausePlayer(player.Name)
		}
		game.Tick()
		game.SendStateToObserver()
		game.Turn++
	}

	for _, player := range game.Map.Players {
		if player.Health > 0 {
			player.Score += ScoreLastMan
			break
		}
	}

	game.Runner.Log("that's all folks!")
	game.SendScores()
	game.Runner.End()
}
